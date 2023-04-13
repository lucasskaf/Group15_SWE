import { Component } from '@angular/core';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { FormBuilder } from '@angular/forms';
import { Movie } from '../user-auth/user';

@Component({
  selector: 'bb-movie-generator',
  templateUrl: './movie-generator.component.html',
  styleUrls: ['./movie-generator.component.css']
})
export class MovieGeneratorComponent {

  generatorForm
  generatedMovie: Movie = {}

  constructor(private movieGeneratorService : MovieGeneratorService,
    private formBuilder: FormBuilder) {
    this.generatorForm = this.formBuilder.group({
      actors: this.formBuilder.control(''),
      genres: this.formBuilder.control(''),
      min_rating: this.formBuilder.control(''),
      max_runtime: this.formBuilder.control(''),
      streaming_providers: this.formBuilder.control('')
    })
  }

  getGeneratedMovie(filters){
    this.movieGeneratorService.getRandomMovieWithFilters(filters).subscribe(
      {
        next: (respMovie) => {
          console.log(respMovie)
        },
        error: (err) => {
          console.log(err)
        }
      }
    )
  }

  onSubmit(generatorData){
    console.log(generatorData)
    generatorData.actors = generatorData.actors.toLowerCase()
    let actors_array = generatorData.actors.split(', ')
    
    generatorData.actors = actors_array
    this.getGeneratedMovie(generatorData)
  }

  formatLabel(value: number): string {
    if(value % 60 == 0){
      return (value / 60) + 'hrs'
    }

    return (Math.floor(value / 60)) + ':' + (value % 60)
  }
}
