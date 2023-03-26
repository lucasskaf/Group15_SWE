import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms'
import { MovieService } from '../../services/movie-service';
import { Movie } from '../../common/movie';

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
      title: new FormControl(''),
      actor: new FormControl(''),
      genre: new FormControl(''),
      rating: new FormControl(''),
      runtime: new FormControl('')
    });
  }

  onSubmit(movie: Movie) {
    console.log(movie);
  }
}
