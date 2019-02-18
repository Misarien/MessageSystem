import { Injectable } from '@angular/core';
import { Protocol } from "./protocol/Protocol";
import { HttpClient } from '@angular/common/http';
import { HttpHeaders }from '@angular/common/http'
import { Observable } from 'rxjs';
import { environment } from '../environments/environment';
import { Long } from 'protobufjs';
import { UserService } from './user.service';
import * as com from './common/im';

@Injectable()
export class WebsocketService {
  
  session_id: string;
  ws: WebSocket;

  global_message: com.GlobalMessage;  //全局全部消息
  nearest_contact:com.NearestContact; // 最近联系人
  address_book: com.AddressBook; // 通讯录
  message_list : com.AllChatRoom; //全部聊天室


  collection: Protocol.Message = new(Protocol.Message);
  constructor(private http:HttpClient,private us:UserService) { 

    this.global_message = new(com.GlobalMessage);
    this.global_message.chat_room_list = new Map<number|Long,com.ChatRoom[]>();

    this.nearest_contact = new(com.NearestContact);
    this.nearest_contact.contact_list = [];

    this.address_book = new(com.AddressBook)
    this.address_book.friends_list = [];
  }


  // 建立websocket链接
  createSocket(url:string){

    let that = this
    this.ws = new WebSocket(url);

    this.ws.onopen = function() {console.log("WebSocket打开");};
    this.ws.onmessage = function(evt) {
      let reader = new FileReader();
      reader.readAsArrayBuffer(evt.data);
      reader.onload = function (e) {
      let buf = new Uint8Array(reader.result as ArrayBuffer);
      let conn = Protocol.Message.decode(buf);
      //that.parseNotification(conn)    //收到消息解析后分析消息
    }};
    this.ws.onclose = function() {console.log("WebSocket关闭")};
  }

  // 请求头部设置x-session-id
  createSessionHeader():HttpHeaders {
    let headers = new HttpHeaders();
    headers = headers.set('X-Session-Id', this.us.session_id);
    console.log("session=", this.us.session_id)
    return headers
  }

  // 获取通讯录
  getAddressBook():Observable<any>{
    let url  = environment.apiUrl+"/address-book"
    return this.http.get(url, {headers:this.createSessionHeader()})
  }

  // 获取最近联系人
  getNearestContact():Observable<any>{
    let url  = environment.apiUrl+"/recent-contact"
    return this.http.get(url, {headers:this.createSessionHeader()})
  }

  //获取最近联系人的最近聊天信息
  getNearestContactMessage(){
    let url  = environment.apiUrl+"/recent-contact-message"
    return this.http.get(url, {headers:this.createSessionHeader(),observe:'response'})
  }
  
  getNearestList(){//获取最近联系人
      this.getNearestContactMessage().subscribe((data) => {
        console.log("最近联系人de消息",data);
        // if(data.contact_list.length == 0){
        //   data.contact_list = [];
        // }
        // let HL =  new(com.NearestContact);
        this.nearest_contact.contact_list = [];
        // this.nearest_contact.contact_list = data.contact_list;
        // console.log("contact_list = ", this.nearest_contact.contact_list)

        // return data
        for(let i=0;i<data.body['chat_room_list'].length;i++){
          let FriItem:com.NearestContactItem = new(com.NearestContactItem);
          FriItem.id=data.body['chat_room_list'][i].id;
          FriItem.name=data.body['chat_room_list'][i].name;
          FriItem.head_img=data.body['chat_room_list'][i].head_img;
          FriItem.is_group=data.body['chat_room_list'][i].is_group;
          FriItem.count=data.body['chat_room_list'][i].count;
          FriItem.message_list = data.body['chat_room_list'][i].message_list;
          this.nearest_contact.contact_list.push(FriItem)
        }
        console.log("contact_list = ", this.nearest_contact.contact_list)
        // this.getNearestMessage();
      })
  }

  // //分析消息
  // parseNotification(conn:Protocol.Message){
  //   if (conn.type==Protocol.Message.Type.NOTIFICATION){
  //     console.log("NOTIFICATION");
  //     if (conn.cmd == Protocol.Message.CtrlType.NONE){
  //       console.log("COMMOM",conn);
  //       for(let i=0;i<this.wsMessageList.List.length;i++){
  //           if(conn.isgroup == this.wsMessageList.List[i].Isgroup&&
  //             ((conn.isgroup &&conn.to==this.wsMessageList.List[i].ID)||(!conn.isgroup &&conn.from==this.wsMessageList.List[i].ID))){
  //             let item = new(MessageItem);
  //             item.Mid = conn.msgid;
  //             item.From = conn.from;
  //             item.To = conn.from;
  //             if(conn.isgroup){
  //               item.To = conn.to;
  //             }
  //             item.Content = conn.content;
  //             item.ContentType =conn.contentType;
  //             item.Time = conn.time;
  //             this.wsMessageList.List[i].MList.push(item);
  //             console.log(this.wsMessageList.List[i]);
  //             break;
  //           }
  //       }
  //     }else if(conn.cmd == Protocol.Message.CtrlType.CREATE_SESSION){ //添加好友请求
  //       this.createSessById(conn,conn.from)
  //     }else if (conn.cmd == Protocol.Message.CtrlType.CREATE_GROUP || conn.cmd == Protocol.Message.CtrlType.GROUP_ADDMEMBERS){
  //       this.createGroupById(conn,conn.to)
  //     }else if(conn.cmd == Protocol.Message.CtrlType.MSG_BACK){
  //       for(let i=0;i<this.wsMessageList.List.length;i++){
  //         if(conn.isgroup == this.wsMessageList.List[i].Isgroup&&
  //           ((conn.isgroup &&conn.to==this.wsMessageList.List[i].ID)||(!conn.isgroup &&conn.from==this.wsMessageList.List[i].ID))){
  //            for(let j=0;j<this.wsMessageList.List[i].MList.length;j++){
  //             if (this.wsMessageList.List[i].MList[j].Mid == conn.msgid){
  //               this.wsMessageList.List[i].MList.slice(j,1);
  //               break;
  //             }
  //           }
  //         }
  //       }
  //     }
  //   }else if(conn.type==Protocol.Message.Type.ACK){
  //     if (conn.cmd == Protocol.Message.CtrlType.NONE){
  //       for(let i=0;i<this.wsMessageList.List.length;i++){
  //         if (conn.from==this.wsMessageList.List[i].ID && conn.isgroup == this.wsMessageList.List[i].Isgroup){
  //           for(let j=0;j<this.wsMessageList.List[i].MList.length;j++){
  //             if (this.wsMessageList.List[i].MList[j].Time == conn.time){
  //               this.wsMessageList.List[i].MList[j].Mid = conn.msgid;
  //               break;
  //             }
  //           }
  //           break;
  //         }
  //       }
  //     }else if(conn.cmd == Protocol.Message.CtrlType.CREATE_SESSION){//添加好友请求确认信息
  //      this.createSessById(conn,conn.to)
  //     }else if (conn.cmd == Protocol.Message.CtrlType.CREATE_GROUP || conn.cmd == Protocol.Message.CtrlType.GROUP_ADDMEMBERS){
  //       this.createGroupById(conn,conn.to)
  //     }
  //   }
  // }


  // createSessById(conn: Protocol.Message,uid: number|Long){
  //       let item = new(FriendItem);
  //       item.ID = uid;
  //       this.us.getUserbyId(item.ID).subscribe((data)=>{
  //         console.log("data=", data)
  //         item.Name = data["Name"];
  //         item.Headimg = data["Img_url"];
  //       });
  //       item.Counter = 1;
  //       item.Isgroup = false;
  //       this.wsFriendList.List.push(item);

  //       let sess = new(Session)
  //       sess.ID = item.ID
  //       sess.Isgroup = false;
  //       sess.MList = [];
  //       this.wsMessageList.List.push(sess);
  // }
  //  createGroupById(conn: Protocol.Message,gid: number|Long){
  //   let item = new(FriendItem);
  //   item.ID = gid;
  //   this.us.getGroupById(item.ID).subscribe((data)=>{
  //     item.Name = data["Name"];
  //     item.Headimg = data["Headimg"];
  //   });
  //   item.Counter = 1;
  //   item.Isgroup = true;
  //   this.wsFriendList.List.push(item);

  //   let sess = new(Session)
  //   sess.ID = item.ID
  //   sess.Isgroup = true;
  //   sess.MList = [];
  //   this.wsMessageList.List.push(sess);
  // }
  

  // // 发送信息，不在这里构造消息体
  // sendMessage(message: Protocol.Message){
  //   message.time = Date.now()
  //   // console.log("mes.contentype=",message.contentType);
  //   if (message.type ==  Protocol.Message.Type.REQUEST) {
  //     if(message.cmd == Protocol.Message.CtrlType.NONE){  // 单聊或群聊
  //       for(let i=0;i<this.wsMessageList.List.length;i++){
  //         if(message.isgroup == this.wsMessageList.List[i].Isgroup&&message.to==this.wsMessageList.List[i].ID){
  //           let item = new(MessageItem);
  //           item.Mid = 0;
  //           item.From = message.from;
  //           item.To = message.to;
  //           item.Content = message.content;
  //           item.ContentType =message.contentType;
  //           item.Time = message.time;
  //           this.wsMessageList.List[i].MList.push(item);
  //           break;
  //         }
  //       }
  //     }
  //   }else if(message.cmd == Protocol.Message.CtrlType.MSG_BACK){   // 撤回消息
  //     if(message.msgid == 0){
  //       alert("消息ＩＤ不存在，无法撤回");
  //     }
  //     for(let i=0;i<this.wsMessageList.List.length;i++){
  //       if (message.to==this.wsMessageList.List[i].ID && message.isgroup == this.wsMessageList.List[i].Isgroup){
  //         for(let j=0;j<this.wsMessageList.List[i].MList.length;j++){
  //           if (this.wsMessageList.List[i].MList[j].Mid == message.msgid){
  //             this.wsMessageList.List[i].MList.slice(j,1);
  //             break;
  //           }
  //         }
  //       }
  //     }
  //   }
  //   console.log("发送前的数据",message)
  //   this.ws.send(Protocol.Message.encode(message).finish());
  // }

  // delectMessage(id: number){//撤回消息

  // }
  // //  下面函数是登录的时候初始数据 
  // //
  // // 获取聊天列表，或者是好友聊天列表
  // getChatList(){
  //   let url  = environment.apiUrl+"/chatlist"
  //   return this.http.get(url)
  // }
  // HistoryMessage(info){
  //   this.getChatMessageList(info).subscribe((data:MsgList) => {
  //     console.log("历史消息原数据",data)
  //     this.wsMessageList.List =[];
  //     if (data.List == null) {
  //       data.List = [];
  //     }
  //     for(let i=0;i<data.List.length;i++){
  //         let session = new(Session);
  //         session.MList = [];
  //         if (data.List[i].length>0){
  //           session.ID = data.List[i][0].From;
  //           if (data.List[i][0].Isgroup ||session.ID == this.us.MyUserId) {
  //             session.ID = data.List[i][0].To;
  //           }
  //           session.Isgroup = data.List[i][0].Isgroup;
  //         }else{continue;}
  //         // console.log("session.ID",session.ID, data.List[i][0].To,data.List[i][0].From)
  //        for(let j=0;j<data.List[i].length;j++){
  //         let Item = new(MessageItem);
  //         Item.Mid = data.List[i][j].Mid;
  //         Item.From = data.List[i][j].From;
  //         Item.To = data.List[i][j].To;
  //         Item.Content = data.List[i][j].Content;
  //         Item.ContentType = data.List[i][j].ContentType;
  //         Item.Time = data.List[i][j].Time;
  //         session.MList.push(Item);
  //       }
  //       this.wsMessageList.List.push(session)
  //       // console.log("ContenType = ", session.MList);
  //     }
  //     console.log("历史消息",this.wsMessageList.List)
  //   })
  // }
  // InitChatList(){
  //   this.getAddressBook().subscribe((data) => {
  //     console.log("通讯录",data);
  //     let HL =  new(com.AddressBook);
  //     HL.friends_list = [];
    //   for(let i=0;i<data.List.length;i++){
    //     let FriItem:FriendItem = new(FriendItem);
    //     FriItem.ID=data.List[i].Id;
    //     FriItem.Name=data.List[i].Name;
    //     FriItem.Headimg=data.List[i].Headimg;
    //     FriItem.Isgroup=data.List[i].Isgroup;
    //     FriItem.Counter=data.List[i].Counter;
    //     this.wsFriendList.List.push(FriItem)

    //     let Ht = new(Hist);
    //     Ht.ID = data.List[i].Id;
    //     Ht.Isgroup = data.List[i].Isgroup;
    //     HL.List.push(Ht)
    //   }
    //   this.HistoryMessage(HL)
  //   })
  // }
}
