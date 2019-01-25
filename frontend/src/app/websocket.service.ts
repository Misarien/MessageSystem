import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Protocol } from './protocol/c2g'

@Injectable()
export class WebsocketService {
ws: WebSocket;
  constructor() { }

  createObservableSocket(url:string):Observable<any>{
    this.ws = new WebSocket(url);
    return new Observable(observer => {
      this.ws.onmessage = (event)=> observer.next(event.data);
      this.ws.onerror = (event) => observer.error(event);
      this.ws.onclose=(event)=>observer.complete();
      
    }) 
  }
  
  sendMessage(message: any){
    this.ws.send(JSON.stringify(message));
  }
}
