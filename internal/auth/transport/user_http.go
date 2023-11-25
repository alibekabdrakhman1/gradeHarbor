package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

type UserHttpTransport struct {
	config config.UserHttpTransport
	logger *zap.SugaredLogger
}

func NewUserHttpTransport(config config.UserHttpTransport, logger *zap.SugaredLogger) *UserHttpTransport {
	return &UserHttpTransport{
		config: config,
		logger: logger,
	}
}

type GetUserResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (ut *UserHttpTransport) GetUser(ctx context.Context, email string) (*GetUserResponse, error) {
	var response *GetUserResponse

	responseBody, err := ut.makeRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/user/v1/user/getByEmail/%s", email),
		ut.config.Timeout,
	)
	if err != nil {
		ut.logger.Errorf("http: makeRequest err: %v", err)
		return nil, fmt.Errorf("failed to makeRequest err: %w", err)
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		ut.logger.Errorf("http: failed to unmarshall response err %v:", err)
		return nil, fmt.Errorf("failed to unmarshall response err: %w", err)
	}
	fmt.Println(response, err)

	return response, nil
}

func (ut *UserHttpTransport) makeRequest(
	ctx context.Context,
	httpMethod string,
	endpoint string,
	timeout time.Duration,
) (b []byte, err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	requestURL := ut.config.Host + endpoint

	req, err := http.NewRequestWithContext(ctx, httpMethod, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to NewRequestWithContext err: %w", err)
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client making http request err: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body err: %w", err)
	}

	return body, nil
}
