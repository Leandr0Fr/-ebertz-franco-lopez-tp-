create or replace function enviar_email_reserva_o_cancelacion() returns trigger as $$
declare
	turno_selec record;
	paciente_selec record;
	asunto text;
	cuerpo text;
begin
	if new.estado != 'cancelado' and new.estado != 'reservado' then
		return new;
	end if;

	if new.estado = 'reservado' then
		asunto := 'Reserva de turno';
		cuerpo := 'Turno reservado. ';
	else
		asunto = 'Cancelación de turno';
		cuerpo := 'Turno cancelado. ';
	end if;
	
	select * into turno_selec from turno where turno.nro_turno = new.nro_turno;
	select * into paciente_selec from paciente where paciente.nro_paciente = turno_selec.nro_paciente; 
	
	if found then
		cuerpo := cuerpo || informacion_turno(turno_selec.nro_turno);
	
		insert into envio_email values (default, current_timestamp,
			paciente_selec.email, asunto, cuerpo, current_timestamp, 'pendiente');
	end if;
	return new;
end;
$$ language plpgsql;

create or replace trigger enviar_email_reserva_o_cancelacion_trg
after update on turno
for each row
execute procedure enviar_email_reserva_o_cancelacion();


create or replace function generar_recordatorios() returns void as $$
declare
	turno_actual record;
	fecha_actual date;
	email_paciente text;
	cuerpo text;
begin
	fecha_actual := current_date;
	
	for turno_actual in select * from turno t where t.estado = 'reservado' and fecha_actual + '2 days'::interval = date_trunc('day', t.fecha) loop
		select email into email_paciente from paciente p where turno_actual.nro_paciente = p.nro_paciente;

		cuerpo := 'Recordatorio de turno reservado. ' || informacion_turno(turno_actual.nro_turno);

		insert into envio_email values (default, current_timestamp,
			email_paciente, 'Recordatorio de turno', cuerpo, current_timestamp, 'pendiente');
	end loop;
end;
$$ language plpgsql;


create or replace function generar_avisos_perdida_de_turnos() returns void as $$
declare
	turno_actual record;
	fecha_actual date;
	email_paciente text;
	cuerpo text;
begin
	fecha_actual := current_date;
	
	for turno_actual in select * from turno t where t.estado = 'reservado' and fecha_actual = date_trunc('day', t.fecha) loop
		select email into email_paciente from paciente p where turno_actual.nro_paciente = p.nro_paciente;
		
		cuerpo := 'Se perdió un turno reservado. ' || informacion_turno(turno_actual.nro_turno);
		
		insert into envio_email values (default, current_timestamp,
			email_paciente, 'Pérdida de turno reservado', cuerpo, current_timestamp, 'pendiente');
	end loop;
end;
$$ language plpgsql;


create or replace function informacion_turno(nro_turno_elegido int) returns text as $$
declare
	turno_elegido record;
	medique_elegido record;
	cuerpo text;
begin
	select * into turno_elegido from turno t where t.nro_turno = nro_turno_elegido;
	select * into medique_elegido from medique m where turno_elegido.dni_medique = m.dni_medique;
	
	cuerpo := 'Turno número ' || nro_turno_elegido ||  ' para la fecha y hora ' || 
	turno_elegido.fecha || ' en el consultorio ' || turno_elegido.nro_consultorio || '. ';
	
	cuerpo := cuerpo || 'Médique asignade: ' || medique_elegido.apellido || ', ' || medique_elegido.nombre ||
	'. Especialidad: ' || medique_elegido.especialidad; 
	
	return cuerpo;
end;
$$ language plpgsql; 
