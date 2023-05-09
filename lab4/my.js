const cities = [
    { name: "New York", population: 8399000, country: "USA" },
    { name: "Virginia", population: 8982000, country: "USA" },
    { name: "Los Angeles", population: 13929286, country: "USA" },
  ];
  
  const citiesDiv = document.getElementById("cities");
  
  let citiesHtml = "";
  
  for (let i = 0; i < cities.length; i++) {
    const city = cities[i];
    const cityHtml = `
      <div class="city">
        <h2>${city.name}</h2>
        <p>Population: ${city.population.toLocaleString()}</p>
        <p>Country: ${city.country}</p>
      </div>
    `;
    citiesHtml += cityHtml;
  }
  
  citiesDiv.innerHTML = citiesHtml;

const numbers = [8399000, 8982000, 13929286];

const sumsDiv = document.getElementById("sums");

const sum = numbers.reduce((acc, curr) => acc + curr, 0);

const average = sum / numbers.length;

const sumsHtml = `
  <p>Sum: ${sum}</p>
  <p>Average: ${average}</p>
`;

sumsDiv.innerHTML = sumsHtml;
