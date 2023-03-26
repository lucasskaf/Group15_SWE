import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-movie',
  templateUrl: './movie.component.html',
  styleUrls: ['./movie.component.css']
})
export class MovieComponent {

  @Input() title!: string;
  @Input() actor!: string;
  @Input() genre!: string;
  @Input() rating!: string;
  @Input() runtime!: string;

  @Input() movieItem!: {
    title: string;
    actor: string;
    genre: string;
    rating: string;
    runtime: string;
  }
}
