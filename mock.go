package main

var mf []food = []food{
	{
		Id:          "94301",
		Name:        "Chicken Livers and Portuguese Roll",
		Position:    2,
		Description: "Chicken Livers Topped with PERi-PERi Sauce, Served with A Roll To Soak Up The Sauce",
		Images:      []string{},
		SubItems: []subItems{
			{
				Id:            "224474",
				Name:          "Chicken Livers and Portuguese Roll",
				Position:      1,
				Price:         "250.00",
				Consumable:    "1:1",
				Cuisine_name:  "Indian",
				Category_name: "Appeteasers",
				Discount: discount{
					Type:   "",
					Amount: "0.00",
				},
				Tags: []string{},
			},
		},
	},
	{
		Id:          "94304",
		Name:        "Spicy Mixed Olives (V)",
		Position:    3,
		Description: "Co-Starring Garlic, Pepper and Chili",
		Images:      []string{},
		SubItems: []subItems{
			{
				Id:            "224477",
				Name:          "Spicy Mixed Olives (V)",
				Position:      1,
				Price:         "215.00",
				Consumable:    "1:1",
				Cuisine_name:  "Indian",
				Category_name: "Appeteasers",
				Discount: discount{
					Type:   "",
					Amount: "0.00",
				},
				Tags: []string{},
			},
		},
	},
}
