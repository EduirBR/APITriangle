function language(lang) {
  form = document.getElementById("form");
  if (lang === "english") {
    form.innerHTML = `<div align="center">
    <div>
        API Draw a Triangle
    </div>
    <table>
        <tr>
            <td size="20"><span>Base Size</span>

            </td>
            <td>
                <input type="text" name="size" size="1" maxlength="2">
            </td>
        </tr>
    </table>
    <button align="center" type='button' onclick=Triangle(form.size.value)>Submit</button>`;
  }

  if (lang === "spanish") {
    form.innerHTML = `<div align="center">
    <div>
        API Dibuja el Triangulo
    </div>
    <table>
        <tr>
            <td size="20"><span>Tamaño de la base</span>

            </td>
            <td>
                <input type="text" name="size" size="1" maxlength="2">
            </td>
        </tr>
    </table>
    <button align="center" type='button' onclick=Triangle(form.size.value)>Submit</button>`;
  }
}

const Triangle = async (size) => {
  size = parseInt(size);
  if (Par(size)) {
    let data = {
      size: parseInt(size),
    };
    try {
      let response = await fetch("http://localhost:8000/triangle", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        const Jdata = await response.json();
        console.log(Jdata);
        let inner = "";
        for (let lines of Jdata) {
          for (let char of lines) {
            if (char == " ") {
              inner += "&nbsp&nbsp";
            } else {
              inner += char;
            }
          }
          inner += "<br>";
        }
        document.getElementById("Result").innerHTML = `<p>${inner}</p>`;
        form.reset();
      } else {
        alert("Error With API");
      }
    } catch (error) {
      alert("Error WIth API");
      console.log(error);
    }
  } else {
    alert("El numero debe ser impar\nNumber must be pair");
  }
};
let Par = (num) => {
  //Verifica si es impar
  //check if is not a pair number
  if (num % 2 != 0) {
    return true;
  } else {
    return false;
  }
};
