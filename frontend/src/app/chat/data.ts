// export class MesList {
//     from: number = 0;
//     to: number = 0;
//     msgid: number = 0;
//     content: string = "";
//     img_url: string = "";
//   }

// //好友列表
// export class FriendList{
//     List: FriendItem[];
// }
// export class FriendItem {
//     ID :number = 0;
//     NAME: string = "";
//     HEADIMG: string = "";
//     ISGROUP: boolean = false;
//     Counter: number = 0;
// }

// //聊天信息
// export class MessageList{
//     List: Session[];
//  }
// export class  Session{
//     ID: number;
//     ISGROUP: boolean;
//     Message: MessageItem[]
// }
// export class MessageItem{
//     Mid: number = 0;
//     From: number = 0;
//     To: number = 0;
//     Gid: number = 0;//群组TD
//     Content: string = "";
//     ContentType: number = 0;
//     Time: number = 0;
// } 
import * as com from '../common/im';

export const addressList : com.AddressBookItem[] = [
    {id: 1, name: "abc", head_img: "aaa", is_group: false},
    {id: 2, name: "bcd", head_img: "aaa", is_group: false},
    {id: 3, name: "cde", head_img: "aaa", is_group: false},
    {id: 4, name: "def", head_img: "aaa", is_group: true},
    {id: 5, name: "efg", head_img: "aaa", is_group: true},
    {id: 6, name: "fgh", head_img: "aaa", is_group: true},
]

export const nearContractList : com.NearestContactItem[] = [
    {id: 2, name: "bcd", head_img: "aaa", count: 5, is_group: false},
    {id: 1, name: "abc", head_img: "aaa", count: 1, is_group: false},
    {id: 3, name: "cde", head_img: "aaa", count: 12, is_group: false},
    {id: 4, name: "def", head_img: "aaa", count: 6, is_group: true},
    {id: 5, name: "efg", head_img: "aaa", count: 2, is_group: true},
    {id: 6, name: "fgh", head_img: "aaa", count: 2, is_group: true},

]

export const id_1_message : com.MessageItem[] = [
    {id: 1001, from: 100018, to: 1, content: "测试", content_type: 0, arrive_time: 1000001, is_group: false},
    {id: 1002, from: 100018, to: 1, content: "测试", content_type: 0, arrive_time: 1000002, is_group: false},
    {id: 1003, from: 1, to: 100018, content: "测试", content_type: 0, arrive_time: 1000003, is_group: false},
    {id: 1004, from: 100018, to: 1, content: "测试", content_type: 0, arrive_time: 1000004, is_group: false},
    {id: 1005, from: 1, to: 100018, content: "测试", content_type: 0, arrive_time: 1000005, is_group: false},
    {id: 1006, from: 1, to: 100018, content: "测试", content_type: 0, arrive_time: 1000006, is_group: false},
    {id: 1007, from: 100018, to: 1, content: "测试", content_type: 0, arrive_time: 1000007, is_group: false},
    {id: 1008, from: 100018, to: 1, content: "测试", content_type: 0, arrive_time: 1000008, is_group: false},
]
export const id_2_message : com.MessageItem[] = [
    {id: 1001, from: 100018, to: 2, content: "测试", content_type: 0, arrive_time: 1000001, is_group: false},
    {id: 1002, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000002, is_group: false},
    {id: 1003, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000003, is_group: false},
    {id: 1004, from: 100018, to: 2, content: "测试", content_type: 0, arrive_time: 1000004, is_group: false},
    {id: 1005, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000005, is_group: false},
    {id: 1006, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000006, is_group: false},
    {id: 1007, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000007, is_group: false},
    {id: 1008, from: 100018, to: 2, content: "测试", content_type: 0, arrive_time: 1000008, is_group: false},
    {id: 1009, from: 100018, to: 2, content: "测试", content_type: 0, arrive_time: 1000004, is_group: false},
    {id: 1010, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000005, is_group: false},
    {id: 1011, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000006, is_group: false},
    {id: 1011, from: 2, to: 100018, content: "测试", content_type: 0, arrive_time: 1000007, is_group: false},
    {id: 1011, from: 100018, to: 2, content: "测试", content_type: 0, arrive_time: 1000008, is_group: false},
]
export const id_3_message : com.MessageItem[] = [
    {id: 1001, from: 3, to: 100018, content: "测试", content_type: 0, arrive_time: 1000001, is_group: false},
    {id: 1001, from: 100018, to: 3, content: "测试", content_type: 0, arrive_time: 1000002, is_group: false},
    {id: 1001, from: 3, to: 100018, content: "测试", content_type: 0, arrive_time: 1000003, is_group: false},
    {id: 1001, from: 100018, to: 3, content: "测试", content_type: 0, arrive_time: 1000004, is_group: false},
    {id: 1001, from: 100018, to: 3, content: "测试", content_type: 0, arrive_time: 1000005, is_group: false},
    {id: 1001, from: 3, to: 100018, content: "测试", content_type: 0, arrive_time: 1000006, is_group: false},
]

export const chatRoom : com.ChatRoom[] = [
    { id: 1, name: "abc", is_group: false, message_list: id_1_message},
    { id: 2, name: "bcd", is_group: false, message_list: id_2_message},
    { id: 3, name: "cde", is_group: false, message_list: id_3_message},
    { id: 1, name: "abc", is_group: false, message_list: id_1_message},
    { id: 1, name: "abc", is_group: false, message_list: id_1_message},
    { id: 1, name: "abc", is_group: false, message_list: id_1_message}

]
