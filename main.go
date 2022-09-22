// Урок https://www.youtube.com/watch?v=wWZ5u-Gyioc&list=PLgG7lPwNdp57Dx2WxQZmXzzcp1Cqxk-e_&index=4&ab_channel=BRO-IT

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(300, 400))

	label1 := widget.NewLabel("Enter the first number")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Enter the second number")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")

	answercheck := widget.NewLabel("")
	answerradio := widget.NewLabel("")

	btn1 := widget.NewButton("+", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)
		if err1 != nil || err2 != nil {
			answer.SetText("Please, enter the number!")
		} else {
			sum := n1 + n2

			answer.SetText(fmt.Sprintf("(+) %.2f", sum))
		}
	})

	btn2 := widget.NewButton("-", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)
		if err1 != nil || err2 != nil {
			answer.SetText("Please, enter the number!")
		} else {
			sub := n1 - n2

			answer.SetText(fmt.Sprintf("(-) %.2f", sub))
		}
	})

	btn3 := widget.NewButton("*", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)
		if err1 != nil || err2 != nil {
			answer.SetText("Please, enter the number!")
		} else {
			mul := n1 * n2

			answer.SetText(fmt.Sprintf("(*) %.2f", mul))
		}
	})

	btn4 := widget.NewButton("/", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)
		if err1 != nil || err2 != nil {
			answer.SetText("Please, enter the number!")
		} else {
			div := n1 / n2

			answer.SetText(fmt.Sprintf("(/) %.2f", div))
		}
	})

	checks := widget.NewCheckGroup([]string{"Check1", "Check2", "Check3"}, func(s []string) {
	})

	radio := widget.NewRadioGroup([]string{"Radio1", "Radio2", "Radio3"}, func(s string) {
	})

	btn5 := widget.NewButton("ChechBox", func() {
		for _, i := range checks.Selected {
			//fmt.Println(i)
			answercheck.SetText(fmt.Sprintf(i))
		}
		fmt.Println(radio.Selected)
		answerradio.SetText(fmt.Sprintf(radio.Selected))
	})

	w.SetContent(container.NewVBox(
		label1, entry1,
		label2, entry2,
		btn1, btn2, btn3, btn4,
		answer,
		checks,
		radio,
		btn5,
		answercheck,
		answerradio,
	))

	w.ShowAndRun()

}
