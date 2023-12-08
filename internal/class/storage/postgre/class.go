package postgre

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/jmoiron/sqlx"
)

func NewClassRepository(db *sqlx.DB) *ClassRepository {
	return &ClassRepository{
		DB: db,
	}
}

type ClassRepository struct {
	DB *sqlx.DB
}

func (r *ClassRepository) GetClassesForTeacher(ctx context.Context, userID uint) ([]*model.Class, error) {
	var classes []*model.Class
	query := `SELECT * FROM class WHERE teacher_id = $1 ORDER BY class_code`
	err := r.DB.SelectContext(ctx, &classes, query, userID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *ClassRepository) GetClassesForStudent(ctx context.Context, userID uint) ([]*model.Class, error) {
	query := `
        SELECT
            c.id AS class_id,
            c.class_code,
            c.class_name,
            c.description,
            c.teacher_id,
            c.teacher_name
        FROM
            class c
        JOIN
            student s ON c.id = s.class_id
        WHERE
            s.student_id = $1
        ORDER BY 
            c.class_code;
    `

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*model.Class
	for rows.Next() {
		var class model.Class
		if err := rows.Scan(
			&class.ID,
			&class.ClassCode,
			&class.ClassName,
			&class.Description,
			&class.TeacherID,
			&class.TeacherName,
		); err != nil {
			return nil, err
		}
		classes = append(classes, &class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepository) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	var res model.ClassWithID
	var class model.Class
	classQuery := "SELECT * FROM class WHERE id = $1"
	err := r.DB.GetContext(ctx, &class, classQuery, id)
	if err != nil {
		return nil, err
	}

	var students []model.ClassStudent
	studentsQuery := "SELECT student_id, student_name FROM student WHERE class_id = $1"
	err = r.DB.SelectContext(ctx, &students, studentsQuery, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(students)
	fmt.Println(class)
	res.ID = class.ID
	res.ClassName = class.ClassName
	res.ClassCode = class.ClassCode
	res.Teacher = model.User{
		ID:       class.TeacherID,
		FullName: class.TeacherName,
	}
	for _, student := range students {
		res.Students = append(res.Students, model.User{
			ID:       student.StudentID,
			FullName: student.StudentName,
		})
	}

	return &res, nil
}

func (r *ClassRepository) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	var students []model.ClassStudent
	classQuery := "SELECT student_id, student_name FROM student WHERE class_id = $1"
	err := r.DB.SelectContext(ctx, &students, classQuery, id)
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

func (r *ClassRepository) GetClassGradesByIDForStudent(ctx context.Context, id uint, userID uint) (*model.Grade, error) {
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
    g.student_id = $1
    AND 
    c.id = $2
ORDER BY 
    g.week;
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

	if err := r.DB.SelectContext(ctx, &grades, query, userID, id); err != nil {
		return nil, err
	}

	classGrades := &model.Grade{}
	if classGrades.ClassCode == "" {
		classGrades.ClassID = grades[0].ClassID
		classGrades.ClassCode = grades[0].ClassCode
		classGrades.ClassName = grades[0].ClassName
		classGrades.Teacher = model.User{
			ID:       grades[0].TeacherID,
			FullName: grades[0].TeacherName,
		}
	}
	classGrades.Students = append(classGrades.Students, model.GradeStudent{
		ID:       grades[0].StudentID,
		FullName: grades[0].StudentName,
		Grades:   []model.GradeResponse{},
	})
	for _, row := range grades {
		if row.Grade.Valid {
			classGrades.Students[0].Grades = append(classGrades.Students[0].Grades, model.GradeResponse{
				Grade:        int(row.Grade.Int64),
				Week:         int(row.Week.Int64),
				LastModified: row.LastModified,
			})
		}
	}

	return classGrades, nil
}

func (r *ClassRepository) GetClassGradesByIDForTeacher(ctx context.Context, id uint, userID uint) (*model.Grade, error) {
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
    c.id = $1
ORDER BY
    student_id, g.week;
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
	fmt.Println(grades)

	if grades[0].TeacherID != userID {
		return nil, errors.New("another teacher id")
	}

	classGrades := &model.Grade{}

	for _, row := range grades {
		if classGrades.ClassCode == "" {
			classGrades.ClassID = grades[0].ClassID
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
		classGrades.Students = append(classGrades.Students, *student)
	}

	return classGrades, nil
}

func (r *ClassRepository) PutClassGradesByID(ctx context.Context, id uint, grades model.GradesRequest) error {
	for _, grade := range grades.Grades {
		query := `
			INSERT INTO grade (class_id, student_id, grade, week, last_modified)
			VALUES ($1, $2, $3, $4, $5)
		`

		_, err := r.DB.ExecContext(ctx, query, id, grade.StudentID, grade.Grade, grade.Week, time.Now())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ClassRepository) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	var class model.User
	fmt.Println(id, "----")
	classQuery := "SELECT teacher_id as id, teacher_name as full_name FROM class WHERE id = $1"
	err := r.DB.GetContext(ctx, &class, classQuery, id)
	fmt.Println(class, "_---")
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       class.ID,
		FullName: class.FullName,
	}, nil
}

func (r *ClassRepository) GetStudentGradesByID(ctx context.Context, studentID uint) ([]*model.Grade, error) {
	classes, err := r.GetClassesForStudent(ctx, studentID)
	fmt.Println(classes[0])
	if err != nil {
		return nil, err
	}
	var res []*model.Grade

	for _, class := range classes {
		grades, err := r.GetClassGradesByIDForStudent(ctx, class.ID, studentID)
		fmt.Println(grades)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.Grade{
			ClassID:   class.ID,
			ClassCode: class.ClassCode,
			ClassName: class.ClassName,
			Teacher: model.User{
				ID:       class.TeacherID,
				FullName: class.TeacherName,
			},
			Students: grades.Students,
		})
	}
	return res, nil
}

func (r *ClassRepository) GetMyStudentsForTeacher(ctx context.Context, id uint) ([]uint, error) {
	var ids []uint
	query := `
		SELECT r.student_id
		FROM relationships r
		WHERE r.teacher_id = $1
	`

	if err := r.DB.SelectContext(ctx, &ids, query, id); err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *ClassRepository) GetMyStudentsForStudent(ctx context.Context, id uint) ([]uint, error) {
	var ids []uint
	query := `
		SELECT s.student_id
		FROM student s
		WHERE s.class_id = (
    		SELECT class_id
    		FROM student
    		WHERE student_id = $1
		)
  		AND s.student_id != $1;
	`

	if err := r.DB.SelectContext(ctx, &ids, query, id); err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *ClassRepository) GetMyTeachers(ctx context.Context, id uint) ([]uint, error) {
	var ids []uint
	query := `
		SELECT teacher_id
		FROM relationships
		WHERE student_id = $1;
	`
	fmt.Println(ids)

	if err := r.DB.SelectContext(ctx, &ids, query, id); err != nil {
		return nil, err
	}
	fmt.Println("0000")

	return ids, nil
}
