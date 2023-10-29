m = require("../gojs-cvss.js")

var vec = m.cvss40.New()
e = vec.Parse("CVSS:4.0/AV:N/AC:L/AT:N/PR:N/UI:N/VC:H/VI:H/VA:H/SC:H/SI:H/SA:H/MVI:L/MSA:S");
if (e != null) {
    console.error(e);
}
if (vec.Get("MAV") == "X") {
    vec.Set("MAV", "A")
}
score = vec.Score();
console.log(vec.Vector()+": "+ score + " "+ Rating(score)[0]);
