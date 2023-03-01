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