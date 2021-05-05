// User

//     - id
//     - username
//     - roomlist
//     - connection
package models;

import (
	"room"
	"connection"
)

type User struct {
	id int;
	name string;
	roomList room.Room[];
	connection connection.Connection;
}