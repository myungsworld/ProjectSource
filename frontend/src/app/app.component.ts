import { Component } from '@angular/core';
import { HEROES } from './mock-heroes';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = "myungsworld"
  heroes = HEROES;
}
