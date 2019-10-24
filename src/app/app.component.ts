import { Component } from '@angular/core';
import { HelloService } from './hello.service';

@Component({
  selector: 'app-root',
  template: `
  {{ message }}
    `
})
export class AppComponent {
  title = 'my-app';
  message: string;
  constructor(private helloService: HelloService) { }

  ngOnInit() {
    this.hello();
    this.hello2();
    this.hello();
  }

  hello(): void {
    this.message = this.helloService.getHello();
  }

  hello2(): void {
    this.helloService.getHelloDelayed()
      .subscribe(message => {
        this.message = message;
       });
  }
}
