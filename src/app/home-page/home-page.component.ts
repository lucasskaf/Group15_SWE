import { HttpClient } from '@angular/common/http';
import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { LoginRegisterService } from '../services/login-register.service';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import {MatAutocompleteSelectedEvent} from '@angular/material/autocomplete';
import { Observable, map, startWith } from 'rxjs';
import {COMMA, ENTER} from '@angular/cdk/keycodes';
import { MatChipInputEvent } from '@angular/material/chips';
import { MatIconModule } from '@angular/material/icon'; 
import {MatSliderModule} from '@angular/material/slider';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {
  isLoginOpen = false
  isAuthenticated
  username = ""
  message = 'Home Page'

  generatorForm

  constructor(
    private formBuilder: FormBuilder
  ){
    this.generatorForm = this.formBuilder.group({
      actorCtrl: this.formBuilder.control(''),
      genresCtrl: this.formBuilder.control(''),
      ratingCtrl: this.formBuilder.control(''),
      runtimeCtrl: this.formBuilder.control(''),
      providerCtrl: this.formBuilder.control('')
    })
  }

  ngOnInit(): void {
    Emmiters.userData.subscribe(
      {
        next: (username : string) => {
          this.username = username
        }
      }
    )

    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
          this.message = 'Hey '
        }
      }
    )

    // this.generatorForm = this.formBuilder.group({
    //   actors: this.actorCtrl,
    //   genre: this.formBuilder.control(''),
    //   rating: this.formBuilder.control(''),
    //   runtime: this.formBuilder.control(''),
    //   provider: this.formBuilder.control('')
    // })
    // this.generatorForm = new FormGroup({
    //   // actors: this.actorCtrl,
    //   genres: new FormControl(''),
    //   minRating: new FormControl(''),
    //   maxRuntime: new FormControl(''),
    //   provider: new FormControl('')
    // })
  }

  onSubmit(generatorData){
    console.log(generatorData)
  }

  formatLabel(value: number): string {
    if(value % 60 == 0){
      return (value / 60) + 'hrs'
    }

    return (Math.floor(value / 60)) + ':' + (value % 60)
  }
}
