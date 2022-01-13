var tmp = 0;

var lt = /</g, 
    gt = />/g, 
    ap = /'/g, 
    ic = /"/g;
value = value.toString().replace(lt, "&lt;").replace(gt, "&gt;").replace(ap, "&#39;").replace(ic, "&#34;");

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