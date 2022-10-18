package main

import "github.com/gofiber/fiber/v2"

func GetIngredients(ctx *fiber.Ctx) string {
	ingstr := ""
	if ctx.FormValue("Ingredient0") != "" {
		ingstr += ctx.FormValue("Ingredient0") + "|"
	}
	if ctx.FormValue("Ingredient1") != "" {
		ingstr += ctx.FormValue("Ingredient1") + "|"
	}
	if ctx.FormValue("Ingredient2") != "" {
		ingstr += ctx.FormValue("Ingredient2") + "|"
	}
	if ctx.FormValue("Ingredient3") != "" {
		ingstr += ctx.FormValue("Ingredient3") + "|"
	}
	if ctx.FormValue("Ingredient4") != "" {
		ingstr += ctx.FormValue("Ingredient4") + "|"
	}
	if ctx.FormValue("Ingredient5") != "" {
		ingstr += ctx.FormValue("Ingredient5") + "|"
	}
	if ctx.FormValue("Ingredient6") != "" {
		ingstr += ctx.FormValue("Ingredient6") + "|"
	}
	outstr := "ChoiceSelfMade" + "|-|" + ctx.FormValue("TypeoDel") + " :Type of Delivery" + "|-|" + ingstr + "Ingridients used in Pizza" + "|-|" + ctx.FormValue("cooked") + " :How much it is cooked"
	return outstr

}

func GetPremadePizza(ctx *fiber.Ctx) string {
	outstr := "ChoicePreMade" + "|-|" + ctx.FormValue("TypeoDel") + " :Type of Delivery" + "|-|" + ctx.FormValue("PreMade") + " :TypeofPizza" + "|-|" + ctx.FormValue("cooked") + " :How much it is cooked"
	return outstr
}

func Formin(ctx *fiber.Ctx) error {
	if ctx.FormValue("chooseTypeoPizza") == "SelfMadePizza" {
		return ctx.SendString(GetIngredients(ctx))
	} else if ctx.FormValue("chooseTypeoPizza") == "PreMadePizza" {
		return ctx.SendString(GetPremadePizza(ctx))
	}
	return ctx.SendString("fail")
}
