// import { Component } from '@angular/core';

// @Component({
//   selector: 'app-root',
//   templateUrl: './app.component.html',
//   styleUrls: ['./app.component.css']
// })
// export class AppComponent {
//   title = 'BingeBuddy';
// }

import {Component, OnInit} from '@angular/core';
import { MovieTestService } from './movie-test.service';
import { Movie } from './movie';
// import { Observable } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
// export class AppComponent implements OnInit {
//   public moviesList: Movie[] = []

//   constructor(private movieService: MovieTestService) {}

//   ngOnInit() {
//    this.movieService.getTitle().subscribe(data => this.moviesList = data);
//   }

// }
export class AppComponent {
  title = 'BingeBuddy';
}
