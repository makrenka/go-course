package practices

import (
	"fmt"
	"slices"
	"strings"
)

func countFriends(m map[string][]string) map[string]int {
	res := make(map[string]int)
	for key, val := range m {
		res[key] = len(val)
	}
	return res
}

func commonFriends(m map[string][]string, user1, user2 string) string {
	compareMap := make(map[string]bool)
	for _, v := range m[user1] {
		compareMap[v] = true
	}

	commonFriends := make([]string, 0)
	for _, v := range m[user2] {
		if compareMap[v] {
			commonFriends = append(commonFriends, v)
			delete(compareMap, v)
		}
	}

	if len(commonFriends) == 0 {
		return "нет общих друзей"
	}

	slices.Sort(commonFriends)

	return strings.Join(commonFriends, ", ")
}

func mostPopularUsers(m map[string]int) (users string, max int) {
	usersSlice := make([]string, 0)
	max = 0
	for _, val := range m {
		if val > max {
			max = val
		}
	}

	for key, val := range m {
		if val == max {
			usersSlice = append(usersSlice, key)
		}
	}
	slices.Sort(usersSlice)

	return strings.Join(usersSlice, ", "), max
}

func PrintResult() {
	friendsData := map[string][]string{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Елена":    {"Алексей", "Сергей", "Дмитрий"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}

	friendsNumber := countFriends(friendsData)
	mostPopularUsers, max := mostPopularUsers(friendsNumber)
	commonFriends := commonFriends(friendsData, "Иван", "Елена")

	fmt.Println("Количество друзей:")

	users := make([]string, 0)
	for key := range friendsNumber {
		users = append(users, key)
	}
	slices.Sort(users)
	for _, v := range users {
		fmt.Printf("%s: %d\n", v, friendsNumber[v])
	}

	fmt.Printf("Общие друзья между пользователями Иван и Елена: %s.\n", commonFriends)

	fmt.Printf("Наиболее популярные пользователи: %s (количество друзей: %d).\n", mostPopularUsers, max)
}
