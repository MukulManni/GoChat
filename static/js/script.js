var tmp = 0;

var lt = /</g, 
    gt = />/g, 
    ap = /'/g, 
    ic = /"/g;

function repeat() {
  fetch("https://go-chatters.herokuapp.com/u/msglist").then(res => res.json()).then(d => {

    if (d.length > tmp) {
      tmp = d.length;

      document.getElementById("globalmsgs").innerText = "";

      for (let i = 0; i < d.length; i++) {
        d[i].msg = d[i].msg.toString().replace(lt, "&lt;").replace(gt, "&gt;").replace(ap, "&#39;").replace(ic, "&#34;");
        d[i].user = d[i].user.toString().replace(lt, "&lt;").replace(gt, "&gt;").replace(ap, "&#39;").replace(ic, "&#34;");
        d[i].color = d[i].color.toString().replace(lt, "&lt;").replace(gt, "&gt;").replace(ap, "&#39;").replace(ic, "&#34;");
        d[i].time = d[i].time.toString().replace(lt, "&lt;").replace(gt, "&gt;").replace(ap, "&#39;").replace(ic, "&#34;");
        
        document.getElementById("globalmsgs").innerHTML += "<div class='msg'><b><font color=" + d[i].color + ">" + d[i].time + d[i].user + ": </font></b>" + d[i].msg + "</div>";
      }
    }
  });
}

tmpinterval = setInterval(repeat, 3000);