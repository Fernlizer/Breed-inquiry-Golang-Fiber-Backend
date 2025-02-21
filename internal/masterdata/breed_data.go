package masterdata

import (
	"time"

	"github.com/Fernlizer/Breed-inquiry-Golang-Fiber-Backend/internal/domain"
)

// BreedMasterData ข้อมูลตั้งต้นของพันธุ์วัว
var BreedMasterData = []domain.Breed{
	{ID: "10011001", NameEn: "ABERDEEN ANGUS", NameTh: "อาเบอร์ดีน แองกัส", ShortName: "AN", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011007", NameEn: "AMERICAN BRAHMAN", NameTh: "อเมริกัน บราห์มัน", ShortName: "ABM", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011040", NameEn: "BEEFMASTER", NameTh: "บีฟมาสเตอร์", ShortName: "BM", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011053", NameEn: "BLACK ANGUS", NameTh: "แบล็คแองกัส", ShortName: "BAN", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011066", NameEn: "BRAHMAN", NameTh: "บราห์มัน", ShortName: "BR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011070", NameEn: "BRANGUS", NameTh: "แบรงกัส", ShortName: "BN", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011086", NameEn: "CHARBRAY X ZEBU", NameTh: "ชาร์เบรย์ x ซีบู", ShortName: "CB", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011088", NameEn: "CHAROLAIS", NameTh: "ชาร์โรเลส์", ShortName: "CH", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011115", NameEn: "DROUGHTMASTER", NameTh: "เดร้าท์มาสเตอร์", ShortName: "DZ", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011162", NameEn: "INDO BRAZILIAN", NameTh: "ฮินดูบราซิล", ShortName: "IB", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011171", NameEn: "KABIN BURI", NameTh: "กบินทร์บุรี", ShortName: "KB", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011187", NameEn: "LIMOUSIN", NameTh: "ลิมูซิน", ShortName: "LM", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011247", NameEn: "RED ANGUS", NameTh: "แองกัสแดง", ShortName: "AR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011248", NameEn: "RED BRAHMAN", NameTh: "บราห์มันแดง", ShortName: "RR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011249", NameEn: "RED BRANGUS", NameTh: "แบรงกัสแดง", ShortName: "RBR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011278", NameEn: "SHORTHORN", NameTh: "ชอตฮอร์น", ShortName: "SS", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011281", NameEn: "SIMMENTAL", NameTh: "ซิมเมนทอล", ShortName: "SM", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011291", NameEn: "TAK", NameTh: "ตาก", ShortName: "TK", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011297", NameEn: "THAI BRAHMAN", NameTh: "ไทยบราห์มัน", ShortName: "TBR", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: "10011315", NameEn: "ZEBU (UNSPECIFIED)", NameTh: "ซีบู", ShortName: "ZE", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}
