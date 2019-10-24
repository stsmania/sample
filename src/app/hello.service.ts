import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HelloService {
  message: string = 'non Observable message';
  delay_message: Subject<string> = new Subject();

  getHello(): string {
    return this.message;
  }

  getHelloDelayed(): Observable<string> {
    let message;
    // return Observable.create(observer=>{
    //   setTimeout(()=>{
    //     observer.next("これはObservable");
    //   },10000);
    // })
     setInterval(() => {
       message = 'Observable message';
       this.delay_message.next(message)
     }, 10000)
    return this.delay_message;
  }
}
