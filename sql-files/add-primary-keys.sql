alter table paciente add constraint paciente_pk primary key (nro_paciente);
alter table medique add constraint medique_pk primary key (dni_medique);
alter table consultorio add constraint consultorio_pk primary key (nro_consultorio);
alter table agenda add constraint agenda_pk primary key (dni_medique, dia);
alter table turno add constraint turno_pk primary key (nro_turno);
alter table reprogramacion add constraint reprogramacion_pk primary key (nro_turno);
alter table error add constraint error_pk primary key (nro_error);
alter table cobertura add constraint cobertura_pk primary key (dni_medique, nro_obra_social);
alter table obra_social add constraint obra_social_pk primary key (nro_obra_social);
alter table liquidacion_cabecera add constraint liquidacion_cabecera_pk primary key (nro_liquidacion);
alter table liquidacion_detalle add constraint liquidacion_detalle_pk primary key (nro_liquidacion, nro_linea);
alter table envio_email add constraint envio_email_pk primary key (nro_email);

-- PK de la tabla que NO es parte del modelo de datos
alter table solicitud_reservas add constraint solicitud_reservas_pk primary key (nro_orden);
