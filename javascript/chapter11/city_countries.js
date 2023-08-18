/**
 * @param {string} city
 * @param {string} country
 * @param {number=} population
 * @returns {string}
 */
export function cityCountry(city, country, population = 0) {
  return `${city}, ${country} - population ${population}`;
}
