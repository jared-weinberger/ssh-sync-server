package dto

import "github.com/google/uuid"

type DataDto struct {
	ID        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Keys      []KeyDto       `json:"keys"`
	MasterKey []byte         `json:"master_key"`
	SshConfig []SshConfigDto `json:"ssh_config"`
}

type KeyDto struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	Filename string    `json:"filename"`
	Data     []byte    `json:"data"`
}

type SshConfigDto struct {
	Host   string            `json:"host"`
	Values map[string]string `json:"values"`
}

type UserDto struct {
	Username string `json:"username"`
}

type UserMachineDto struct {
	Username    string `json:"username"`
	MachineName string `json:"machine_name"`
}
