import { Component } from '@angular/core';
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
  testMovie: Movie = {
    adult: false,
    backdrop_path: '/892rHLop6XobpTmdJ05UfkWMZpf.jpg',
    id: 789708,
    original_language: 'en',
    original_title: 'Hilda and the Mountain King',
    overview:
      'When Hilda wakes up in the body of a troll, she must use her wits and courage to get back home, become human again — and save the city of Trolberg.',
    popularity: 36.58,
    poster_path: '/Ac9xJvb1oTnv71j2yfdCQktIlYT.jpg',
    release_date: '2021-12-30',
    title: 'Hilda and the Mountain King',
    vote_average: 7.5,
    vote_count: 114,
  };

  username;
  generatorForm;
  postForm;
  generatedMovie: Movie = { vote_average: 0.0 };
  openPopup: boolean = false;

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
  }

  ngOnInit(): void {
    Emmiters.isPopupOpen.subscribe({
      next: (status) => {
        this.openPopup = status;
      },
    });
  }

  getGeneratedMovie(filters) {
    this.movieGeneratorService.getRandomMovieWithFilters(filters).subscribe({
      next: (respMovie) => {
        console.log(respMovie);
        this.generatedMovie = respMovie;
        Emmiters.isPopupOpen.emit(true);
      },
      error: (err) => {
        console.log(err);
      },
    });
  }

  onSubmit(generatorData) {
    console.log(`BFORE: ${generatorData}`);

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

    console.log(`AFTER: ${generatorData}`);

    this.getGeneratedMovie(generatorData);
  }

  formatLabel(value: number): string {
    if (value % 60 == 0) {
      return value / 60 + 'hrs';
    }

    return Math.floor(value / 60) + ':' + (value % 60);
  }
}
