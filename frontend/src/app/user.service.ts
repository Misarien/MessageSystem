import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';
import { HttpHeaders }from '@angular/common/http';

@Injectable()
export class UserService {
  constructor(private http: HttpClient) { }
  public session_id = '';
  public isLogin = false;
  public MyUserId = -1;
  public myImg = "";
  public myName = ""
  configUrl = environment.apiUrl;
  loginUrl = this.configUrl+"/login"
  signupUrl = this.configUrl+'/register'
  quitUrl = this.configUrl+'/quit'
  postLoginData(data) {
    return this.http.post(this.loginUrl, data,{observe:'response'});
  }
  postSignupData(data) {
    return this.http.post(this.signupUrl, data,{observe:'response'});
  }
  quit(){

    return this.http.get(this.quitUrl,{headers:this.createSessionHeader()});
  }
  getUserbyId(id){
    let url = this.configUrl+'/user?id='+id
    return this.http.get(url)
  }
  getGroupById(id){
    let url = this.configUrl+'/group?id='+id
    let headers = new Headers();
    //headers.append('X-Session-Id','')
    return this.http.get(url)
  }
  getuserinfo(id){
    let url = this.configUrl+'/user-info/'+id
    return this.http.get(url)
  }
  getgroupinfo(id){
    let url = this.configUrl+'/group-info/'+id
    return this.http.get(url)
  }
  createSessionHeader():HttpHeaders {
    let headers = new HttpHeaders();
    headers = headers.set('X-Session-Id', this.session_id);
    console.log("session=", this.session_id)
    return headers
  }
}
