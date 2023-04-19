import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {

  @Output() isClosed = new EventEmitter<boolean>();
  @Input() isOpen : boolean = false;
  isAddMoviePopupOpen: boolean = false;
  isAddPostPopupOpen: boolean = false;

  constructor() { }

  ngOnInit(): void {

  }

  public toggleAddMoviePopupStatus(event: boolean): void {
    this.isAddMoviePopupOpen = event;
  }

  public toggleAddPostPopupStatus(event: boolean): void {
    this.isAddPostPopupOpen = event;
  }
}