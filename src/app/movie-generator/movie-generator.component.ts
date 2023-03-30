import { Component } from '@angular/core';
import { MovieGeneratorService } from '../services/movie-generator.service';

@Component({
  selector: 'bb-movie-generator',
  templateUrl: './movie-generator.component.html',
  styleUrls: ['./movie-generator.component.css']
})
export class MovieGeneratorComponent {

  constructor(private movieGeneratorService : MovieGeneratorService) {}

  movie

  getGeneratedMovie(){
    this.movieGeneratorService.getMovie().subscribe((movie) => {
      console.log(movie)
      this.movie = movie
    })
  }
}
