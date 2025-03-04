package lbcli

import "github.com/fatih/color"

var COLOR_LOG = color.RGB(128, 128, 128)

func Log(str string) {
	COLOR_LOG.Println("[i]\t" + str)
}

func Inform(str string) {
	color.HiCyan("[I]\t" + str)
}

func Warn(str string) {
	color.Yellow("[!]\t" + str)
}

func Fatal(str string) {
	color.HiRed("[!]\t" + str)
}

func Action(str string) {
	color.HiMagenta("[*]\t" + str)
}
