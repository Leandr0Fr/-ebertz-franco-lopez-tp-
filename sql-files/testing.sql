drop database if exists turnos;
create database turnos;

\c turnos

\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/create-tables.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/add-primary-keys.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/add-foreign-keys.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/add-data.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/generacion-de-turnos.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/reserva-de-turnos.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/cancelacion-de-turnos.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/atencion-de-turno.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/liquidacion-obras-sociales.sql'
\i '/home/lilo/ebertz-franco-lopez-tp/sql-files/generar-emails.sql'

select generar_turnos_disponibles(2023, 06);

-- dni de medique no valido
insert into solicitud_reservas values (1, 1, 0, '2023-06-15', '15:00:00');

-- nro de historia clinica no válido
insert into solicitud_reservas values (2, 0, 24730016, '2023-06-19', '8:45:00');

-- paciente con obra social no atendida por medique
insert into solicitud_reservas values (3, 5, 41724061, '2023-06-21', '12:30:00');

-- turno inexistente 
insert into solicitud_reservas values (4, 3, 29732435, '2000-01-01', '00:00:00');

-- turno disponible
insert into solicitud_reservas values (5, 11, 49194249, '2023-06-28', '11:15:00');

-- turno no disponible
insert into solicitud_reservas values (6, 11, 49194249, '2023-06-28', '11:15:00');

-- paciente sin obra social con turno disponible
insert into solicitud_reservas values (7, 21, 21376991, '2023-06-23', '08:00:00');

-- supera limite de reserva de turnos
insert into solicitud_reservas values (8, 1, 24730016, '2023-06-23', '10:30:00');
insert into solicitud_reservas values (9, 1, 18715278, '2023-06-15', '09:45:00');
insert into solicitud_reservas values (10, 1, 24730016, '2023-06-30', '09:00:00');
insert into solicitud_reservas values (11, 1, 18715278, '2023-06-22', '11:15:00');
insert into solicitud_reservas values (12, 1, 24730016, '2023-06-30', '11:15:00');
insert into solicitud_reservas values (13, 1, 18715278, '2023-06-29', '11:15:00');

select reservar_turnos();

select count(*) emails_turnos_reservados from envio_email where asunto = 'Reserva de turno'; -- 7

select cancelar_turnos(18715278, '2023-06-01', current_date - '1 day'::interval);
select cancelar_turnos(49194249, '2023-06-01', '2023-06-28');

select count(*) emails_turnos_cancelados from envio_email where asunto = 'Cancelación de turno'; -- 117

-- nro de turno no valido
select atender_paciente(0);

-- turno no reservado
select atender_paciente(1);

create or replace function reserva_random_para_fecha(fecha_pedida date) returns int as $$
declare
	datos_turno record;
	nro_obra_social_buscada int;
	nro_paciente_buscado int;
begin
	if extract(dow from fecha_pedida) = 0 then -- Si es domingo
		return null;
	end if;
	select * into datos_turno from turno where date(fecha) = fecha_pedida and estado = 'disponible';
	select nro_obra_social into nro_obra_social_buscada from cobertura where dni_medique = datos_turno.dni_medique;
	select nro_paciente into nro_paciente_buscado from paciente p1 where nro_obra_social = nro_obra_social_buscada and not exists (
		select 1 from turno t where t.nro_paciente = p1.nro_paciente and estado = 'reservado' group by t.nro_paciente having count(*) >= 5
	);
	perform reservar_turno(nro_paciente_buscado, datos_turno.dni_medique, date(datos_turno.fecha), to_char(datos_turno.fecha, 'hh:mi:ss')::time);
	return datos_turno.nro_turno;
end;
$$ language plpgsql;

-- el turno no corresponde con el dia actual. Nota: los domingos no se puede reservar
select atender_paciente(reserva_random_para_fecha(date(current_date + '1 day'::interval)));

-- el turno corresponde con el dia actual. Nota: los domingos no se puede reservar
select atender_paciente(reserva_random_para_fecha(current_date));

--select liquidar_obras_sociales(2023, 06);

select reserva_random_para_fecha(date(current_date + '2 day'::interval));
select generar_recordatorios();

select reserva_random_para_fecha(current_date);
select generar_avisos_perdida_de_turnos();
