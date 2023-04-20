import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators, ValidationErrors } from '@angular/forms'
import { MovieService } from '../../services/movie-service';
import { MovieComponent } from '../../common/movie/movie.component';

@Component({
  selector: 'app-movie-form',
  templateUrl: './movie-form.component.html',
  styleUrls: ['./movie-form.component.css']
})
export class MovieFormComponent implements OnInit {
  form!: FormGroup;

  constructor(private movieService: MovieService) { }

  ngOnInit() {
    this.form = new FormGroup({
      title: new FormControl('', Validators.compose([
        Validators.required,
        Validators.pattern('[\\w\\-\\s\\/]+')
      ])),
      actor: new FormControl('', Validators.compose([
        Validators.required,
        Validators.pattern('[a-zA-Z\\s]+')
      ])),
      genre: new FormControl('', Validators.compose([
        Validators.required
      ])),
      rating: new FormControl('', Validators.compose([
        Validators.required,
        this.ratingValidator
      ])),
      runtime: new FormControl('', Validators.compose([
        Validators.required,
        this.runtimeValidator
      ]))
    });
  }

  // Only allows for singe digit values between 1-10.
  ratingValidator(control: FormControl): ValidationErrors | null {
    const rating = control.value;

    if (!rating || !/^[1-9]$|^10$/.test(rating)) {
      return { rating: { min: 1, max: 10 } };
    }

    return null;
  }

  // Only allows for runtimes below 999.
  runtimeValidator(control: FormControl): ValidationErrors | null  {
    const runtime = control.value;
    const maxRuntime = 999;
    const minRuntime = 1;

    if (!runtime || parseInt(runtime, 10) > maxRuntime) {
      return {
        runtime: {
          min: minRuntime,
          max: maxRuntime
        }
      };
    }
    return null;
  }

  onSubmit(movie: MovieComponent) {
    // console.log(movie);
    this.movieService.addWatchedMovie(movie);
    this.form.reset();
  }
}
