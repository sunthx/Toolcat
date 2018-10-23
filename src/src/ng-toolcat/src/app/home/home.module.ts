import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { HomeRoutes } from './home.routes';
import { GenerateGuidComponent } from '../generate-guid/generate-guid.component';
import { HomeComponent } from '../home/home.component';
import { GenerateQrComponent } from '../generate-qr/generate-qr.component';
import { TimeConverterComponent } from '../time-converter/time-converter.component';
import { EncryptionComponent } from '../encryption/encryption.component';

@NgModule({
    declarations: [
        GenerateGuidComponent,
        HomeComponent,
        GenerateQrComponent,
        TimeConverterComponent,
        EncryptionComponent
    ],
    imports: [
        RouterModule.forChild(HomeRoutes)
    ],
    providers: [],
    bootstrap: [HomeComponent]
})
export class HomeModule { }
