alter table paciente drop constraint if exists paciente_nro_obra_social_fk;

alter table agenda drop constraint if exists agenda_dni_medique_fk;
alter table agenda drop constraint if exists agenda_nro_consultorio_fk;

alter table turno drop constraint if exists turno_nro_consultorio_fk;
alter table turno drop constraint if exists turno_dni_medique_fk;
alter table turno drop constraint if exists turno_nro_paciente_fk;
alter table turno drop constraint if exists turno_nro_obra_social_consulta_fk;

alter table reprogramacion drop constraint if exists reprogramacion_nro_turno_fk;

alter table error drop constraint if exists error_nro_consultorio_fk;

alter table cobertura drop constraint if exists cobertura_dni_medique_fk;
alter table cobertura drop constraint if exists cobertura_nro_obra_social_fk;

alter table liquidacion_cabecera drop constraint if exists liquidacion_cabecera_nro_obra_social_fk;

alter table liquidacion_detalle drop constraint if exists liquidacion_detalle_nro_liquidacion_fk;
alter table liquidacion_detalle drop constraint if exists liquidacion_detalle_dni_medique_fk;

alter table paciente drop constraint if exists paciente_pk;
alter table medique drop constraint if exists medique_pk;
alter table consultorio drop constraint if exists consultorio_pk;
alter table agenda drop constraint if exists agenda_pk;
alter table turno drop constraint if exists turno_pk;
alter table reprogramacion drop constraint if exists reprogramacion_pk;
alter table error drop constraint if exists error_pk;
alter table cobertura drop constraint if exists cobertura_pk;
alter table obra_social drop constraint if exists obra_social_pk;
alter table liquidacion_cabecera drop constraint if exists liquidacion_cabecera_pk;
alter table liquidacion_detalle drop constraint if exists liquidacion_detalle_pk;
alter table envio_email drop constraint if exists envio_email_pk;
alter table solicitud_reservas drop constraint if exists solicitud_reservas_pk;
