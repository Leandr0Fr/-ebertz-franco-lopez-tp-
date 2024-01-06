alter table paciente add constraint paciente_nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table agenda add constraint agenda_dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table agenda add constraint agenda_nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);

alter table turno add constraint turno_nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table turno add constraint turno_dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table turno add constraint turno_nro_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);
alter table turno add constraint turno_nro_obra_social_consulta_fk foreign key (nro_obra_social_consulta) references obra_social (nro_obra_social);

alter table reprogramacion add constraint reprogramacion_nro_turno_fk foreign key (nro_turno) references turno (nro_turno);

alter table error add constraint error_nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);

alter table cobertura add constraint cobertura_dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table cobertura add constraint cobertura_nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table liquidacion_cabecera add constraint liquidacion_cabecera_nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

alter table liquidacion_detalle add constraint liquidacion_detalle_nro_liquidacion_fk foreign key (nro_liquidacion) references liquidacion_cabecera (nro_liquidacion);
alter table liquidacion_detalle add constraint liquidacion_detalle_dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
