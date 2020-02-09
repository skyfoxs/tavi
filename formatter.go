package tavi

import (
	"fmt"
	"time"

	"github.com/leekchan/accounting"
)

func getFormatted(amount int64) string {
	if amount/1000000 > 0 {
		return getFormatted(amount/1000000) + "ล้าน" + getFormatted(amount%1000000)
	}
	if amount/100000 > 0 {
		return getFormatted(amount/100000) + "แสน" + getFormatted(amount%100000)
	}
	if amount/10000 > 0 {
		return getFormatted(amount/10000) + "หมื่น" + getFormatted(amount%10000)
	}
	if amount/1000 > 0 {
		return getFormatted(amount/1000) + "พัน" + getFormatted(amount%1000)
	}
	if amount/100 > 0 {
		return getFormatted(amount/100) + "ร้อย" + getFormatted(amount%100)
	}
	if amount == 21 {
		return "ยี่สิบเอ็ด"
	}
	if amount/20 == 1 {
		return "ยี่สิบ" + getFormatted(amount%20)
	}
	if amount == 11 {
		return "สิบเอ็ด"
	}
	if amount/10 == 1 {
		return "สิบ" + getFormatted(amount%10)
	}
	if amount/10 > 0 {
		if amount%10 == 1 {
			return getFormatted(amount/10) + "สิบเอ็ด"
		}
		return getFormatted(amount/10) + "สิบ" + getFormatted(amount%10)
	}
	text := []string{"", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	return text[amount]
}

func formatMoney(amount int64) string {
	ac := accounting.Accounting{Precision: 0, FormatZero: "00"}
	return ac.FormatMoney(amount)
}

func formatDate(t time.Time) string {
	var months = [...]string{
		"ม.ค.",
		"ก.พ.",
		"มี.ค.",
		"เม.ย.",
		"พ.ค.",
		"มิ.ย.",
		"ก.ค.",
		"ส.ค.",
		"ก.ย.",
		"ต.ค.",
		"พ.ย.",
		"ธ.ค.",
	}
	year, month, day := t.Date()
	return fmt.Sprintf("%d %s %d", day, months[month], year+543)
}
