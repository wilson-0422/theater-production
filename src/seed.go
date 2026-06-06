package main

import (
	"log"
	"theater-production/src/config"
	"theater-production/src/models"
	"theater-production/src/services"
	"time"
)

func Run() {
	db := config.GetDB()

	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	log.Println("正在初始化种子数据...")

	services.CreateUser("admin", "admin123", "admin")
	services.CreateUser("director", "director123", "director")
	services.CreateUser("staff", "staff123", "staff")

	scripts := []models.Script{
		{Title: "雷雨", Author: "曹禺", Genre: "话剧", Synopsis: "《雷雨》是剧作家曹禺创作的一部话剧，发表于1934年。此剧以1925年前后的中国社会为背景，描写了一个带有浓厚封建色彩的资产阶级家庭的悲剧。", Duration: 120, Status: "performing"},
		{Title: "茶馆", Author: "老舍", Genre: "话剧", Synopsis: "《茶馆》是老舍于1956年创作的话剧，以北京一家茶馆为背景，展示了清末、民初、抗战胜利后三个不同时期的社会变迁。", Duration: 150, Status: "rehearsing"},
		{Title: "牡丹亭", Author: "汤显祖", Genre: "昆曲", Synopsis: "《牡丹亭》是明代剧作家汤显祖的代表作，描写了杜丽娘与柳梦梅的爱情故事，是中国戏曲史上最杰出的作品之一。", Duration: 180, Status: "performing"},
		{Title: "龙须沟", Author: "老舍", Genre: "话剧", Synopsis: "《龙须沟》是老舍于1950年创作的话剧，通过描写北京龙须沟旁劳动人民生活的变化，歌颂了新中国。", Duration: 130, Status: "draft"},
		{Title: "日出", Author: "曹禺", Genre: "话剧", Synopsis: "《日出》是曹禺于1935年创作的话剧，以抗战前的天津社会为背景，揭露了半殖民地半封建社会的黑暗现实。", Duration: 110, Status: "archived"},
	}
	for i := range scripts {
		db.Create(&scripts[i])
	}

	actors := []models.Actor{
		{Name: "张明远", Gender: "男", Phone: "13800138001", RoleType: "lead", Skills: "话剧表演、台词功底、武术", Status: "busy"},
		{Name: "李婉清", Gender: "女", Phone: "13800138002", RoleType: "lead", Skills: "昆曲表演、舞蹈、声乐", Status: "busy"},
		{Name: "王建国", Gender: "男", Phone: "13800138003", RoleType: "supporting", Skills: "话剧表演、喜剧表演", Status: "available"},
		{Name: "陈雅芝", Gender: "女", Phone: "13800138004", RoleType: "supporting", Skills: "话剧表演、声乐、舞蹈", Status: "available"},
		{Name: "刘宏伟", Gender: "男", Phone: "13800138005", RoleType: "extra", Skills: "群众演员、武术", Status: "available"},
		{Name: "赵雪梅", Gender: "女", Phone: "13800138006", RoleType: "supporting", Skills: "昆曲表演、古筝", Status: "leave"},
	}
	for i := range actors {
		db.Create(&actors[i])
	}

	props := []models.Prop{
		{Name: "青花瓷瓶", Category: "prop", Description: "清代风格仿古青花瓷瓶，用于茶馆布景", Quantity: 5, Available: 3, Status: "available"},
		{Name: "长衫（男）", Category: "costume", Description: "民国时期男式长衫，藏青色", Quantity: 10, Available: 6, Status: "available"},
		{Name: "旗袍（女）", Category: "costume", Description: "民国时期女式旗袍，多种颜色", Quantity: 8, Available: 5, Status: "available"},
		{Name: "茶具套装", Category: "prop", Description: "老北京茶馆风格茶具，含茶壶、茶碗、茶盘", Quantity: 3, Available: 2, Status: "available"},
		{Name: "古典折扇", Category: "prop", Description: "昆曲表演用折扇，丝绸扇面", Quantity: 15, Available: 12, Status: "available"},
		{Name: "水袖戏服", Category: "costume", Description: "昆曲女角水袖戏服，白色绣花", Quantity: 4, Available: 2, Status: "low"},
	}
	for i := range props {
		db.Create(&props[i])
	}

	theaters := []models.Theater{
		{Name: "国家大剧院", Location: "北京市西城区西长安街2号", Capacity: 3000, Contact: "010-66550000", Facilities: "主舞台、侧舞台、旋转舞台、交响乐池"},
		{Name: "上海大剧院", Location: "上海市黄浦区人民大道300号", Capacity: 1800, Contact: "021-63868686", Facilities: "主舞台、后舞台、双侧台、乐池"},
		{Name: "广州大剧院", Location: "广州市天河区珠江新城", Capacity: 2000, Contact: "020-38392888", Facilities: "主舞台、可升降乐池、多媒体系统"},
	}
	for i := range theaters {
		db.Create(&theaters[i])
	}

	tours := []models.Tour{
		{ScriptID: 1, City: "上海", Venue: "上海大剧院", StartDate: time.Date(2026, 7, 1, 0, 0, 0, 0, time.Local), EndDate: time.Date(2026, 7, 7, 0, 0, 0, 0, time.Local), Status: "confirmed", Notes: "雷雨全国巡演上海站，共7场演出"},
		{ScriptID: 2, City: "广州", Venue: "广州大剧院", StartDate: time.Date(2026, 8, 15, 0, 0, 0, 0, time.Local), EndDate: time.Date(2026, 8, 25, 0, 0, 0, 0, time.Local), Status: "planned", Notes: "茶馆全国巡演广州站，共10场演出"},
		{ScriptID: 3, City: "南京", Venue: "江苏大剧院", StartDate: time.Date(2026, 9, 10, 0, 0, 0, 0, time.Local), EndDate: time.Date(2026, 9, 15, 0, 0, 0, 0, time.Local), Status: "planned", Notes: "牡丹亭全国巡演南京站"},
	}
	for i := range tours {
		db.Create(&tours[i])
	}

	schedules := []models.ActorSchedule{
		{ActorID: 1, ScriptID: 1, RoleName: "周朴园", ScheduleDate: time.Date(2026, 6, 10, 0, 0, 0, 0, time.Local), Status: "confirmed"},
		{ActorID: 2, ScriptID: 3, RoleName: "杜丽娘", ScheduleDate: time.Date(2026, 6, 12, 0, 0, 0, 0, time.Local), Status: "confirmed"},
		{ActorID: 3, ScriptID: 1, RoleName: "鲁贵", ScheduleDate: time.Date(2026, 6, 10, 0, 0, 0, 0, time.Local), Status: "scheduled"},
		{ActorID: 4, ScriptID: 2, RoleName: "王淑芬", ScheduleDate: time.Date(2026, 6, 15, 0, 0, 0, 0, time.Local), Status: "scheduled"},
	}
	for i := range schedules {
		db.Create(&schedules[i])
	}

	requisitions := []models.PropRequisition{
		{PropID: 2, ActorID: 1, Quantity: 2, RequisitionDate: time.Date(2026, 6, 1, 0, 0, 0, 0, time.Local), ReturnDate: time.Date(2026, 7, 10, 0, 0, 0, 0, time.Local), Status: "approved", Notes: "雷雨演出用"},
		{PropID: 4, ActorID: 3, Quantity: 1, RequisitionDate: time.Date(2026, 6, 5, 0, 0, 0, 0, time.Local), Status: "pending", Notes: "茶馆排练用"},
		{PropID: 5, ActorID: 2, Quantity: 1, RequisitionDate: time.Date(2026, 6, 8, 0, 0, 0, 0, time.Local), ReturnDate: time.Date(2026, 9, 20, 0, 0, 0, 0, time.Local), Status: "approved", Notes: "牡丹亭演出用"},
	}
	for i := range requisitions {
		db.Create(&requisitions[i])
	}

	theaterSchedules := []models.TheaterSchedule{
		{TheaterID: 1, ScriptID: 1, StartTime: time.Date(2026, 6, 20, 19, 0, 0, 0, time.Local), EndTime: time.Date(2026, 6, 20, 21, 30, 0, 0, time.Local), Status: "confirmed"},
		{TheaterID: 1, ScriptID: 3, StartTime: time.Date(2026, 6, 25, 19, 0, 0, 0, time.Local), EndTime: time.Date(2026, 6, 25, 22, 0, 0, 0, time.Local), Status: "confirmed"},
		{TheaterID: 2, ScriptID: 2, StartTime: time.Date(2026, 7, 1, 19, 30, 0, 0, time.Local), EndTime: time.Date(2026, 7, 1, 22, 0, 0, 0, time.Local), Status: "booked"},
	}
	for i := range theaterSchedules {
		db.Create(&theaterSchedules[i])
	}

	log.Println("种子数据初始化完成")
}
