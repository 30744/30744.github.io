function funCall(){
  var myArray = [
  {name: "Virginia", population: 8654542, id:1 },
  { name: "Virginia-beach", population: 457672, id:2 },
  { name: "Richmond", population: 226604, id:3 },
  { name: "Norfolk", population: 238832, id:4 }];
  
  var html="<table border='1|1'>"
  setTimeout(()=>{
    for(var i=0; i<myArray.length;i++){
      html+='<tr>';
       html+='<td>'+myArray[i].id+'</td>';
        html+='<td>'+myArray[i].name+'</td>';
        html+='<td>'+myArray[i].population+'</td>';
        html+='</tr>';
    }
    document.getElementById("table").innerHTML=html
      },500); 
    }
      funCall()

      const header = document.querySelector("#header");
const navLinks = document.querySelectorAll(".nav_link");
const tabsBtn = document.querySelectorAll(".tabs__nav-btn");
const tabsItems = document.querySelectorAll(".tabs__item");

tabsBtn.forEach(onTabClick);

function onTabClick(item) {
    item.addEventListener("click", function() {
        console.log("clicked")
            let currentBtn = item;
            let tabId = currentBtn.getAttribute("data-tab");
            let currentTab = document.querySelector(tabId);

        if( ! currentBtn.classList.contains("active")) {
              tabsBtn.forEach(function(item) {
            item.classList.remove("active");
        });

        tabsItems.forEach(function(item) {
            item.classList.remove("active");
        });

        currentBtn.classList.add("active");
        currentTab.classList.add("active");
        }
    });
}
document.querySelector(".tabs__nav-btn").click();



window.addEventListener("scroll", checkScroll);

document.addEventListener("DOMContentLoaded", function() {
    checkScroll();
});

function checkScroll() {
    let scrollPos = window.scrollY;

    if(scrollPos > 0) {
        header.classList.add("green");
        }
    else{
        header.classList.remove("green");
    }
}

testBtn.addEventListener("click", function() {
   console.log("clicked");
});

for(let navItem of navLinks) {
    navItem.addEventListener("click", function() {
   console.log(navItem.text);
});
}

  let person = {
    name: "Gulnaz",
    age: 18,
    address: {
      street: "Doswell St",
      city: "Richmond",
      state: "VG"
    }
  };
  console.log(person.name); 
  console.log(person.address.city); 
  console.log(person["age"]); 

  const sum = [1, 2, 3].reduce(add, 0);
function add(accumulator, a) {
  return accumulator + a;
}
console.log(sum);