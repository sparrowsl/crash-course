import { expect, test } from "vitest";
import { cityCountry } from "./city_countries.js";

test("Format city and country", () => {
  expect(cityCountry("paris", "france", 50)).toBe("paris, france - population 50");
});

test("Format city, country and population", () => {
  expect(cityCountry("paris", "france")).toBe("paris, france - population 0");
  expect(cityCountry("paris", "france", 500)).toBe("paris, france - population 500");
});
