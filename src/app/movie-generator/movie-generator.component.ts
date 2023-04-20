import { Component, Input } from '@angular/core';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { FormBuilder, Validators } from '@angular/forms';
import { Movie, User, moviePosts } from '../user-auth/user';
import { OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';

@Component({
  selector: 'bb-movie-generator',
  templateUrl: './movie-generator.component.html',
  styleUrls: ['./movie-generator.component.css'],
})
export class MovieGeneratorComponent implements OnInit {
  username;
  generatorForm;
  postForm;
  generatedMovie: Movie = {id: 3, vote_average: 0}

  constructor(
    private movieGeneratorService: MovieGeneratorService,
    private formBuilder: FormBuilder
  ) {
    this.generatorForm = this.formBuilder.group({
      actors: this.formBuilder.control(''),
      genres: this.formBuilder.control(''),
      min_rating: this.formBuilder.control(''),
      max_runtime: this.formBuilder.control(''),
      streaming_providers: this.formBuilder.control(''),
    });
<<<<<<< Updated upstream
=======
  }

  ngOnInit(): void {
    
  }

  getGeneratedMovie(filters) {
    this.movieGeneratorService.getRandomMovieWithFilters(filters).subscribe({
      next: (respMovie) => {
        console.log(respMovie);
        // let testMovie = {
        //     "adult": false,
        //     "backdrop_path": "/OQeVFQIq0UmzStgkgugQGf6uHB.jpg",
        //     "id": 662409,
        //     "original_language": "en",
        //     "original_title": "Wake N Bake by Rohan Joshi",
        //     "overview": "\"Hysterical\", \"unmissable\", \"magnificent\", \"profound\" are all words. Coincidentally, Wake N Bake, Rohan Joshi's first stand-up special, also has words. After almost a decade in comedy, one of India's foremost comedians and online has words to say about life in one's thirties, home renovation, (thrilling, we know), not being cut out for marriage or roadtrips, and living a 420- friendly life (oh so now we have your attention. Typical). Some of those words are even funny. The seamless hour-long narrative is a tour of all the things that keep Rohan up at night, which, as it turns out, is pretty trivial stuff, because he's basic like that. It's an hour of comedy you'll never forget for five minutes.",
        //     "popularity": 2.156,
        //     "poster_path": "/qQEiA6R52nSKYG4MOlN3ZRC1VyC.jpg",
        //     "release_date": "2020-01-10",
        //     "title": "Wake N Bake by Rohan Joshi",
        //     "vote_average": 8,
        //     "vote_count": 3
        // }
        
        // console.log(testMovie)
        // this.generatedMovie = respMovie;
        this.generatedMovie = respMovie
        Emmiters.generatedMovie.emit(respMovie)

        for(let i = 0; i < Emmiters.watchList.length; i++){
          if(Emmiters.watchList.at(i)?.id == this.generatedMovie.id){
            Emmiters.isMovieWatched.emit(true);
            console.log('movie was watched');
          }
        }

        Emmiters.isPopupOpen.emit(true);
      },
      error: (err) => {
        console.log(err);
      },
    });
  }

  onSubmit(generatorData) {
    let actors_array: string[] = [];
    // Text process
    if (generatorData.actors != '') {
      generatorData.actors = generatorData.actors.toLowerCase();
      actors_array = generatorData.actors.split(', ');
    }

    generatorData.actors = actors_array;

    //Rating & Runtime process
    if (generatorData.min_rating) {
      generatorData.min_rating = parseFloat(generatorData.min_rating) * 2.0;
    }

    if (generatorData.max_runtime) {
      generatorData.max_runtime = parseInt(generatorData.max_runtime);
    }

    this.getGeneratedMovie(generatorData);
  }

  formatLabel(value: number): string {
    if (value % 60 == 0) {
      return value / 60 + 'hrs';
    }

    return Math.floor(value / 60) + ':' + (value % 60);
>>>>>>> Stashed changes
  }

  ngOnInit(): void {
    
  }

  getGeneratedMovie(filters) {
    this.movieGeneratorService.getRandomMovieWithFilters(filters).subscribe({
      next: (respMovie) => {
        console.log(respMovie);
        // let testMovie = {
        //     "adult": false,
        //     "backdrop_path": "/OQeVFQIq0UmzStgkgugQGf6uHB.jpg",
        //     "id": 662409,
        //     "original_language": "en",
        //     "original_title": "Wake N Bake by Rohan Joshi",
        //     "overview": "\"Hysterical\", \"unmissable\", \"magnificent\", \"profound\" are all words. Coincidentally, Wake N Bake, Rohan Joshi's first stand-up special, also has words. After almost a decade in comedy, one of India's foremost comedians and online has words to say about life in one's thirties, home renovation, (thrilling, we know), not being cut out for marriage or roadtrips, and living a 420- friendly life (oh so now we have your attention. Typical). Some of those words are even funny. The seamless hour-long narrative is a tour of all the things that keep Rohan up at night, which, as it turns out, is pretty trivial stuff, because he's basic like that. It's an hour of comedy you'll never forget for five minutes.",
        //     "popularity": 2.156,
        //     "poster_path": "/qQEiA6R52nSKYG4MOlN3ZRC1VyC.jpg",
        //     "release_date": "2020-01-10",
        //     "title": "Wake N Bake by Rohan Joshi",
        //     "vote_average": 8,
        //     "vote_count": 3
        // }
        
        // console.log(testMovie)
        // this.generatedMovie = respMovie;
        this.generatedMovie = respMovie
        Emmiters.generatedMovie.emit(respMovie)

        for(let i = 0; i < Emmiters.watchList.length; i++){
          if(Emmiters.watchList.at(i)?.id == this.generatedMovie.id){
            Emmiters.isMovieWatched.emit(true);
            console.log('movie was watched');
          }
        }

        Emmiters.isPopupOpen.emit(true);
      },
      error: (err) => {
        console.log(err);
      },
    });
  }

  onSubmit(generatorData) {
    let actors_array: string[] = [];
    // Text process
    if (generatorData.actors != '') {
      generatorData.actors = generatorData.actors.toLowerCase();
      actors_array = generatorData.actors.split(', ');
    }

    generatorData.actors = actors_array;

    //Rating & Runtime process
    if (generatorData.min_rating) {
      generatorData.min_rating = parseFloat(generatorData.min_rating) * 2.0;
    }

    if (generatorData.max_runtime) {
      generatorData.max_runtime = parseInt(generatorData.max_runtime);
    }

    this.getGeneratedMovie(generatorData);
  }

  formatLabel(value: number): string {
    if (value % 60 == 0) {
      return value / 60 + 'hrs';
    }

    return Math.floor(value / 60) + ':' + (value % 60);
  }
}