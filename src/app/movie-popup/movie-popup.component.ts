import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Movie, moviePosts } from '../user-auth/user';
import { OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-movie-popup',
  templateUrl: './movie-popup.component.html',
  styleUrls: ['./movie-popup.component.css']
})
export class MoviePopupComponent implements OnInit {
  @Input() generatedMovie: Movie = {vote_average: 0.0}
  @Input() isPopupOpen: boolean = false
  @Output() ouputPopupStatus: EventEmitter<boolean> = new EventEmitter();

  postForm
  username: string = ""
  isPostOpen: boolean = false
  didSendReview: boolean = false
  didFoundPosts: boolean = false
  moviePosts: moviePosts[] = []
  isAuthenticated: boolean = false
  isMovieWatched: boolean = false

  constructor(private movieGeneratorService : MovieGeneratorService,
    private formBuilder: FormBuilder) {
      // Emmiters.watchList.subscribe(
      //   {
      //     next: (userList: Movie[]) => {
      //       this.userWatchedList = userList
      //       console.log(`SUBSRIBED AND THIS LENGHT: ${this.userWatchedList.length}`)
      //     }
      //   }
      // )
    }

  ngOnInit(): void {
    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
        }
      }
    )

    Emmiters.userData.subscribe(
      {
        next: (username: string) => {
          this.username = username
        }
      }
    )

    Emmiters.isPopupOpen.subscribe(
      {
        next: (status) => {
          this.isPopupOpen = status
        }
      }
    )

    this.setupMoviePopUp()
  }

  showMovie() {
    Emmiters.isPopupOpen.emit(true)
  }

  closeMovie() {
    console.log('CLICKED ON CLOSE')
    Emmiters.isPopupOpen.emit(false)
    this.ouputPopupStatus.emit(false)
  }

  toogleAddPost(){
    this.isPostOpen = !this.isPostOpen
  }

  setupMoviePopUp(){
    if(Emmiters.watchList.includes(this.generatedMovie)){
      this.isMovieWatched = true
    }
    else {
      this.isMovieWatched = false
    }

    this.generateMoviePosts(this.generatedMovie.id?.toString())

    console.log(`WAS MOVIE WATCHED: ${this.isMovieWatched}`)
    this.showMovie()

    console.log(`GETTING USERNAME: ${this.username}`)

    this.postForm = this.formBuilder.group({
      movie_id: this.generatedMovie.id,
      username: this.username,
      title: this.formBuilder.control(''),
      body: this.formBuilder.control('', Validators.required)
    })
  }

  onSubmitPost(post: moviePosts){
    post.movie_id = this.generatedMovie.id?.toString()
    post.username = this.username
    console.log(post)

    this.movieGeneratorService.addMoviePosts(post).subscribe(
      {
        next: (resp) => {
          this.moviePosts.push(post)
          this.didFoundPosts = true
          this.didSendReview = true
          console.log('ADDED SUCCESSFULLY')
        },
        error: (err) => {
          console.log('NOT ADD POST')
        }
      }
    )
  }

  addToWatchlist(movie: Movie) {
    this.movieGeneratorService.addToWatchList(movie).subscribe(
      {
        next: (resp) => {
          console.log(resp)
          this.isMovieWatched = true
          console.log(`BEFORE ADDED: ${Emmiters.watchList.length}`)

          Emmiters.watchList.push(resp)
          console.log(`AFTER ADDED: ${Emmiters.watchList.length}`)
        },
        error: (err) => {
          console.log(err)
        }
      }
    )
  }

  generateMoviePosts(movieID: string | undefined){
    this.movieGeneratorService.getMoviePosts(movieID).subscribe(
      {
        next: (resp) => {
          console.log(`POSTS: ${resp}`)
          this.moviePosts = resp
          this.didFoundPosts = true
        },
        error: (err) => {
          console.log('ERROR IN GETTING POSTS')
          if(err.status == 404){
            this.didFoundPosts = false
          }
        }
      }
    )
  }
}
