import { Component, OnInit } from '@angular/core';
import { GuidService } from './services/guid.service';

@Component({
  selector: 'app-generate-guid',
  templateUrl: './generate-guid.component.html'
})
export class GenerateGuidComponent implements OnInit {
  public newGuid: string;
  constructor(private guidService: GuidService) { }

  ngOnInit() {
  }

  public generateGuid(): void {
    this.guidService.new().subscribe((data: object) => {
      this.newGuid = data['Content'];
    });
  }
}
