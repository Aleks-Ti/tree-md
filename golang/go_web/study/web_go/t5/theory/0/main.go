package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID       int64  `json:"id"`    // json:"values" - это алиасы для сериализации
	Email    string `json:"email"` // Так можно указывать, в каком виде должно называться свойство в JSON-строке в зависимости от соглашений в проекте
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"` //можно указать опцию omitempty. Эта опция говорит о том, что если значение свойства является нулевым, то это поле не нужно включать в сериализованную строку
	password string
}

func main() {
	user := User{
		ID:       222,
		Email:    "john@doe.com",
		Name:     "John",
		Age:      25,
		password: "secret",
	}

	// Сериализация структуры в строку
	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		// Используем Fatal только для примера,
		// нельзя использовать в реальных приложениях
		log.Fatalln("marshal ", err.Error())
	}

	fmt.Printf("JSON string: %s\n", string(jsonBytes))

	// Десериализация строки в структуру
	deserializedUser := User{}
	err = json.Unmarshal(jsonBytes, &deserializedUser)
	if err != nil {
		// Используем Fatal только для примера,
		// нельзя использовать в реальных приложениях
		log.Fatalln("unmarshal ", err.Error())
	}

	fmt.Printf("Deserialized user struct: %+v\n", deserializedUser)
}
