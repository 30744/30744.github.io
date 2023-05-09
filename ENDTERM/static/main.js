function AboutMe(person) {
    var interests = person.interests.join(", ");
    var aboutMeText = "Hi, my name is " + person.name + " and I am " + person.age + " years old. ";
    aboutMeText += "Some of my interests include: " + interests + ".";
    return {
      name: person.name,
      aboutMeText: aboutMeText
    };
  }
  
  var person = {
    name: "Gulnaz",
    age: 18,
    interests: ["playing chess", "reading", "singing","traveling"]
  };
  var outputDiv = document.getElementById("output");
  var aboutMe = AboutMe(person);
  outputDiv.innerHTML = aboutMe.aboutMeText;
  
  
  function myBooks() {
      const books = [
        { title: "Jane Eyre", author: "Charlotte Bronte", year: 1847 },
        { title: "The Gadfly", author: "Ethel L.Voynich", year: 1897 },
        { title: "Murder on the Orient Express", author: "Agatha Christie", year: 1934 },
      ];
      let bookList = "";
      for (let i = 0; i < books.length; i++) {
  
        bookList += `<li><b>Title:</b> ${books[i].title} <br><b>Author:</b> ${books[i].author} <br><b>Year:</b> ${books[i].year}</li>`;
      }
    
      document.getElementById("book-list").innerHTML = bookList;
    
      return {
          bookCount: books.length,
          newestBook: books.reduce((a, b) => (a.year > b.year ? a : b)).title,
          oldestBook: books.reduce((a, b) => (a.year < b.year ? a : b)).title,
      };
    }
    
    const bookInfo = myBooks();
    console.log(bookInfo);