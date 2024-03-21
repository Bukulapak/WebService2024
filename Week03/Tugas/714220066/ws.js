
function PostIssue(device,version,username,email,issue){
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    
    var raw = JSON.stringify({
        "device": device,
        "version": version,
        "username": username,
        "email": email,
        "issue": issue
      });
    
    const requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
      redirect: "follow"
    };
    
    fetch("https://eovmipqknhagwn9.m.pipedream.net", requestOptions)
      .then((response) => response.text())
      .then((result) => GetResponse(result))
      .catch((error) => console.error(error));
}

function GetResponse(result){
    document.getElementById("formIssue").innerHTML = result;
}
    
function PushButton(){
    device=document.getElementById("device").value;
    version=document.getElementById("version").value;
    username=document.getElementById("username").value;
    email=document.getElementById("email").value;
    issue=document.getElementById("issue_desc").value;
    PostIssue(device,version,username,email,issue);
}
    
    