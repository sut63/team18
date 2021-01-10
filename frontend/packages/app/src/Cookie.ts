class Cookies {

    SetCookie(name:any,value:any,days:any) {
        var expires = "";
        if (days) {
            var date = new Date();
            date.setTime(date.getTime() + (days*24*60*60*1000));
            expires = "; expires=" + date.toUTCString();
        }
        document.cookie = name + "=" + (value || "")  + expires + "; path=/";
    }
    
    GetCookie(name:any="user_email") {
        var nameEQ = name + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
        }
        return null;
    }

    CheckLogin(arr:Array<any>, email:any, password:any){
        var boo = false
        arr.forEach((item) => {
            if(item.email === email && item.password === password){
                boo = true                
            }
        });
        return boo
    }

    ClearCookie(name:any="user_email",id:any="user_id"){
        console.log("name in ClearCookie => "+name);
        console.log("id in ClearCookie => "+id);
        const date = new Date();
        date.setTime(date.getTime() + (-1 * 24 * 60 * 60 * 1000));
        document.cookie = name+"=; expires="+date.toUTCString()+"; path=/";
        document.cookie = id+"=; expires="+date.toUTCString()+"; path=/";
    }
}
export {Cookies}