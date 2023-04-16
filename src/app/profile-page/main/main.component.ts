import { Component, Output, Input, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent {
  @Output() isClosed = new EventEmitter<boolean>();
  @Input() isOpen : boolean = false;
  isAddMoviePopupOpen: boolean = false;
  isAddPostPopupOpen: boolean = false;

  constructor() { }

  public toggleAddMoviePopupStatus(event: boolean): void {
    this.isAddMoviePopupOpen = event;
  }

  public toggleAddPostPopupStatus(event: boolean): void {
    this.isAddPostPopupOpen = event;
  }
}