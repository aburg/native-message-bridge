package util

import (
	"github.com/aburg/native-message-bridge/settings"
	"github.com/godbus/dbus/v5"
)

func dbusConnect() *dbus.Conn {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	return conn
}

func dbusMsg(conn *dbus.Conn, msg string) {
	if err := conn.Emit(settings.DbusPath, settings.DbusName, msg); err != nil {
		panic(err)
	}
}

func DbusMsg(msg string) {
	conn := dbusConnect()
	dbusMsg(conn, msg)

	defer conn.Close()
}
