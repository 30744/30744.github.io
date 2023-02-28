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

      function loopNumbers() {
        let numbersElement = document.getElementById("numbers");
        let numbersHtml = "";
        for (let i = 1; i <= 4; i++) {
          numbersHtml += i + "<br>";
        }
        numbersElement.innerHTML = numbersHtml;
      }
  
      function addNumbers(a, b, c) {
        return a + b+c;
      }
  
      function createPerson(name, age) {
        let person = {
          name: name,
          age: age
        };
        
        return person;
      }
  
      function updateAboutMe() {
        let resultElement = document.getElementById("result");
        resultElement.innerHTML = addNumbers(8654542, 457672, 238832);
  
        let personElement = document.getElementById("person");
        let person1 = createPerson("Gulnaz", 18);
        personElement.innerHTML = "Name: " + person1.name + "<br>Age: " + person1.age;
  
        loopNumbers();
      }