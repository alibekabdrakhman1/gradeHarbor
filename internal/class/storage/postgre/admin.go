package postgre

import (
	"context"
	"database/sql"
	"time"

	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/jmoiron/sqlx"
)

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

type AdminRepository struct {
	DB *sqlx.DB
}

func (r *AdminRepository) CreateClass(ctx context.Context, class model.ClassRequest) (uint, error) {
	id, err := r.createClass(ctx, &class)
	if err != nil {
		return 0, err
	}
	err = r.createStudents(ctx, class.Students, id)
	if err != nil {
		return 0, err
	}
	err = r.createRelationships(ctx, class.Teacher.ID, class.Students, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AdminRepository) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	var classes []*model.Class
	query := `SELECT * FROM class`
	err := r.DB.SelectContext(ctx, &classes, query)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *AdminRepository) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	var res model.ClassWithID
	var class model.Class
	classQuery := "SELECT * FROM class WHERE id = $1"
	err := r.DB.GetContext(ctx, &class, classQuery, id)
	if err != nil {
		return nil, err
	}
	var students []model.User
	studentsQuery := "SELECT student_id, student_name FROM student WHERE class_id = $1"
	err = r.DB.SelectContext(ctx, &students, studentsQuery, id)
	if err != nil {
		return nil, err
	}

	res.ID = class.ID
	res.ClassName = class.ClassName
	res.ClassCode = class.ClassCode
	res.Teacher = model.User{
		ID:       class.TeacherID,
		FullName: class.TeacherName,
	}
	for _, student := range students {
		res.Students = append(res.Students, model.User{
			ID:       student.ID,
			FullName: student.FullName,
		})
	}

	return &res, nil
}

func (r *AdminRepository) UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error) {
	oldClass, err := r.GetClassByID(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedClass := model.ClassWithID{
		ID:          oldClass.ID,
		ClassCode:   class.ClassCode,
		ClassName:   class.ClassName,
		Description: class.Description,
		Teacher:     class.Teacher,
		Students:    oldClass.Students,
	}

	updateQuery := `
		UPDATE class
		SET class_code = $1, class_name = $2, description = $3, teacher_id = $4
		WHERE id = $5
	`

	_, err = r.DB.ExecContext(ctx, updateQuery,
		updatedClass.ClassCode,
		updatedClass.ClassName,
		updatedClass.Description,
		updatedClass.Teacher.ID,
		updatedClass.ID,
	)
	if err != nil {
		return nil, err
	}

	return &updatedClass, nil
}

func (r *AdminRepository) DeleteClassByID(ctx context.Context, id uint) error {
	query := "DELETE FROM class WHERE id = $1"

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminRepository) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	var students []model.ClassStudent
	classQuery := "SELECT student_id, student_name FROM student WHERE class_id = $1"
	err := r.DB.GetContext(ctx, &students, classQuery, id)
	if err != nil {
		return nil, err
	}

	var response []*model.User

	for _, student := range students {
		response = append(response, &model.User{
			ID:       student.StudentID,
			FullName: student.StudentName,
		})
	}
	return response, nil
}

func (r *AdminRepository) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	query := `SELECT
    c.id AS class_id,
    c.class_code,
    c.class_name,
    c.description,
    c.teacher_id,
    c.teacher_name,
    stu.student_id AS student_id,
    stu.student_name,
    g.id AS grade_id,
    g.grade,
    g.week,
    g.last_modified
FROM
    class c
LEFT JOIN
    student stu ON c.id = stu.class_id
LEFT JOIN
    grade g ON stu.student_id = g.student_id
WHERE
    c.id = $1;
`

	var grades []struct {
		ClassID      uint          `db:"class_id"`
		ClassCode    string        `db:"class_code"`
		ClassName    string        `db:"class_name"`
		Description  string        `db:"description"`
		TeacherID    uint          `db:"teacher_id"`
		TeacherName  string        `db:"teacher_name"`
		StudentID    uint          `db:"student_id"`
		StudentName  string        `db:"student_name"`
		GradeID      uint          `db:"grade_id"`
		Grade        sql.NullInt64 `db:"grade"`
		Week         sql.NullInt64 `db:"week"`
		LastModified time.Time     `db:"last_modified"`
	}

	if err := r.DB.SelectContext(ctx, &grades, query, id); err != nil {
		return nil, err
	}

	classGrades := &model.Grade{}

	for _, row := range grades {
		if classGrades.ClassCode == "" {
			classGrades.ClassCode = row.ClassCode
			classGrades.ClassName = row.ClassName
			classGrades.Teacher = model.User{
				ID:       row.TeacherID,
				FullName: row.TeacherName,
			}
		}

		student := findOrCreateStudent(classGrades.Students, row.StudentID, row.StudentName)
		if row.Grade.Valid {
			student.Grades = append(student.Grades, model.GradeResponse{
				Grade:        int(row.Grade.Int64),
				Week:         int(row.Week.Int64),
				LastModified: row.LastModified,
			})
		}
	}

	return classGrades, nil
}

func (r *AdminRepository) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	var class model.Class
	classQuery := "SELECT teacher_id, teacher_name FROM class WHERE id = $1"
	err := r.DB.GetContext(ctx, &class, classQuery, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       class.TeacherID,
		FullName: class.TeacherName,
	}, nil
}

func (r *AdminRepository) createClass(ctx context.Context, classRequest *model.ClassRequest) (uint, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var classID int
	query := `INSERT INTO class (class_code, class_name, description, teacher_id, teacher_name) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRowxContext(ctx, query, classRequest.ClassCode, classRequest.ClassName, classRequest.Description, classRequest.Teacher.ID, classRequest.Teacher.FullName).Scan(&classID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint(classID), nil
}

func (r *AdminRepository) createStudents(ctx context.Context, students []model.User, classID uint) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, student := range students {
		query := `INSERT INTO student (class_id, student_id, student_name) VALUES ($1, $2, $3)`
		_, err = tx.ExecContext(ctx, query, classID, student.ID, student.FullName)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminRepository) createRelationships(ctx context.Context, teacherID uint, students []model.User, classID uint) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, student := range students {
		query := `INSERT INTO relationships (student_id, teacher_id, class_id) VALUES ($1, $2, $3)`
		_, _ = tx.ExecContext(ctx, query, student.ID, teacherID, classID)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func findOrCreateStudent(students []model.GradeStudent, id uint, fullName string) *model.GradeStudent {
	for _, student := range students {
		if student.ID == id {
			return &student
		}
	}

	newStudent := model.GradeStudent{
		ID:       id,
		FullName: fullName,
	}
	students = append(students, newStudent)
	return &students[len(students)-1]
}
