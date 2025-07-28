package gentoken

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

const jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcmciOiJmN2JlZmQ5NWVkYTk3MTdmNTM2MzE0MGUxNjBhOGU5N2NmMDllOWZhMzEwZjUwYWUxNjM2NTk3NGZjMjA5N2U2IiwidWlkIjoiMTFkNzY5MjczMWY4NDZhMmExMGYyNjc1MGQ5ZDAyNTEiLCJjbGllbnQiOiJjNzU4MzRiOTA0YzBkM2MzOWNkZmEyNWJkMDkxOWFjNzVlYmY4ZDhhOWMyYTgyODc5NWNmZDA4YmEyOWJhMDA5Iiwic2NvcGUiOiJyZWFkLHBhc3Nwb3J0IiwibmJmIjoxNzUzMzQxMzY1LCJpYXQiOjE3NTMzNDEzNjUsImV4cCI6MTc1MzM0ODU2NSwianRpIjoiZTZhY2RjNzItMDM1NS00MDM3LWI4OWUtOTBlNjZjODEyZWY3IiwiaXNzIjoiZG9tYWluLmNvbS5jbiIsInN1YiI6Il9jNzU4MzRiOTA0YzBkM2MzOWNkZmEyNWJkMDkxOWFjNzVlYmY4ZDhhOWMyYTgyODc5NWNmZDA4YmEyOWJhMDA5XzExZDc2OTI3MzFmODQ2YTJhMTBmMjY3NTBkOWQwMjUxIiwiYXVkIjpbImM3NTgzNGI5MDRjMGQzYzM5Y2RmYTI1YmQwOTE5YWM3NWViZjhkOGE5YzJhODI4Nzk1Y2ZkMDhiYTI5YmEwMDkiXX0.qVEw0pukEXobRhqQamThNOwYXwcQhHKwFIQM2RxJ780"

func DecodeJWT(token string) (header, payload map[string]interface{}, err error) {
	// 1. 按 '.' 分割 JWT
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("invalid JWT format: expected 3 parts, got %d", len(parts))
	}

	// 2. 解码 Header
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode header: %v", err)
	}

	// 3. 解码 Payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode payload: %v", err)
	}

	// 4. 解析 JSON
	header = make(map[string]interface{})
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return nil, nil, fmt.Errorf("failed to parse header JSON: %v", err)
	}

	payload = make(map[string]interface{})
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, nil, fmt.Errorf("failed to parse payload JSON: %v", err)
	}

	return header, payload, nil
}

func TestDecode(t *testing.T) {
	// 解码 JWT
	header, payload, err := DecodeJWT(jwtToken)
	if err != nil {
		fmt.Printf("Error decoding JWT: %v\n", err)
		t.Fail()
		return
	}

	// 打印 Header
	fmt.Println("=== Header ===")
	headerJSON, _ := json.MarshalIndent(header, "", "  ")
	fmt.Println(string(headerJSON))

	// 打印 Payload
	fmt.Println("\n=== Payload ===")
	payloadJSON, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(payloadJSON))
}
