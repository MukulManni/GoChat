var tmp = 0;

function repeat() {
  fetch("https://go-chatters.herokuapp.com/u/msglist").then(res => res.json()).then(d => {
    
  if(d.length > tmp){
    tmp = d.length;

    document.getElementById("globalmsgs").innerText = "";
    
    for (let i = 0; i < d.length; i++) {
      document.getElementById("globalmsgs").innerHTML += "<div class='msg'><b><font color="+d[i].color+">" + string(d[i].user) + ": </font></b>" + string(d[i].msg) + "</div>";
    }
  }
  })
}

setInterval(repeat, 3000);
