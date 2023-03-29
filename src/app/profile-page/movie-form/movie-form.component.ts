import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms'
import { MovieService } from '../../services/movie-service';
import { MovieComponent } from '../../common/movie/movie.component';

@Component({
  selector: 'app-movie-form',
  templateUrl: './movie-form.component.html',
  styleUrls: ['./movie-form.component.css']
})
export class MovieFormComponent implements OnInit {
  form!: FormGroup;

  constructor(private movieService: MovieService) {}

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
        this.ratingValidator
      ]))
    });
  }

  ratingValidator (control: FormControl) {
    const rating = control.value;
    const maxRating = 999;

    if (!/^\d{1,3}$/.test(rating) || parseInt(rating, 10) > maxRating) {
      return { rating: true };
    }
    return null;
  }

  onSubmit(movie: MovieComponent) {
    // console.log(movie);
    this.movieService.addWatchedMovie(movie);
    this.form.reset();
  }
}
