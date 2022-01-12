var tmp = 0;

function repeat() {
  fetch("http://localhost:8080/u/msglist").then(res => res.json()).then(d => {

    if (d.length > tmp) {
      tmp = d.length;

      document.getElementById("globalmsgs").innerText = "";

      for (let i = 0; i < d.length; i++) {
        document.getElementById("globalmsgs").innerHTML += "<div class='msg'><b><font color=" + d[i].color + ">" + jsEscape(String(d[i].user)) + ": </font></b>" + jsEscape(String(d[i].msg)) + "</div>";
      }
    }
  });
}

tmpinterval = setInterval(repeat, 3000);

function lol(){
  clearInterval(tmpinterval);
  tmpinterval = setInterval(repeat, 3000);
}

function jsEscape(str) {
  return String(str).replace(/[^\w. ]/gi, function (c) {
    return '\\u' + ('0000' + c.charCodeAt(0).toString(16)).slice(-4);
  });

}