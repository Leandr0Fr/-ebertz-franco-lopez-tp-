create or replace function reservar_turno(nro_historia_clinica int, dni_medique_pedide int, fecha_pedida date, hora_pedida time) returns boolean as $$
declare
	datos_paciente record;
	datos_cobertura record;
	nro_turno_disponible int;
	nro_consultorio_medique int;
	monto_consulta decimal(12, 2); -- para pacientes sin obra social
begin
	select monto_consulta_privada into monto_consulta from medique where dni_medique = dni_medique_pedide;
	
	if not found then
		insert into error values (default, fecha_pedida + hora_pedida, null, dni_medique_pedide, nro_historia_clinica, 
			'reserva', current_timestamp, '?dni de medique no válido');
		return false;
	end if;
	
	select nro_consultorio into nro_consultorio_medique from agenda where dni_medique = dni_medique_pedide and dia = extract(dow from fecha_pedida);
	
	select * into datos_paciente from paciente where nro_paciente = nro_historia_clinica;
	
	if not found then
		insert into error values (default, fecha_pedida + hora_pedida, nro_consultorio_medique, dni_medique_pedide, nro_historia_clinica, 
			'reserva', current_timestamp, '?nro de historia clinica no válido');
		return false;
	end if;
	
	if datos_paciente.nro_obra_social is not null then
	
		select * into datos_cobertura from cobertura where dni_medique = dni_medique_pedide and nro_obra_social = datos_paciente.nro_obra_social;
		
		if not found then
			insert into error values (default, fecha_pedida + hora_pedida, nro_consultorio_medique, dni_medique_pedide, nro_historia_clinica, 
				'reserva', current_timestamp, '?obra social de paciente no atendida por le medique');
			return false;
		end if;
		
	end if;
	
	select nro_turno into nro_turno_disponible from turno 
	where dni_medique = dni_medique_pedide and fecha = fecha_pedida + hora_pedida and estado = 'disponible';
	
	if not found then
		insert into error values (default, fecha_pedida + hora_pedida, nro_consultorio_medique, dni_medique_pedide, nro_historia_clinica, 
			'reserva', current_timestamp, '?turno inexistente o no disponible');
		return false;
	end if;
	
	perform 1 from turno where nro_paciente = nro_historia_clinica and estado = 'reservado' group by nro_paciente having count(*) >= 5;
	
	if found then
		insert into error values (default, fecha_pedida + hora_pedida, nro_consultorio_medique, dni_medique_pedide, nro_historia_clinica, 
			'reserva', current_timestamp, '?supera limite de reserva de turnos');
		return false;
	end if;
	
	-- Turno aprobado
	if datos_paciente.nro_obra_social is null then
		update turno set monto_paciente = monto_consulta where nro_turno = nro_turno_disponible;
	else
		update turno set monto_paciente = datos_cobertura.monto_paciente, monto_obra_social = datos_cobertura.monto_obra_social 
			where nro_turno = nro_turno_disponible;
	end if;
	
	update turno set 
		nro_consultorio = nro_consultorio_medique,
		nro_paciente = nro_historia_clinica,
		nro_obra_social_consulta = datos_paciente.nro_obra_social,
		nro_afiliade_consulta = datos_paciente.nro_afiliade,
		f_reserva = current_timestamp,
		estado = 'reservado' where nro_turno = nro_turno_disponible;
	
	return true;
end;
$$ language plpgsql;

create or replace function reservar_turnos() returns void as $$
declare
	v record;
begin
	for v in select * from solicitud_reservas order by nro_orden loop
		perform reservar_turno(v.nro_paciente, v.dni_medique, v.fecha, v.hora);
	end loop;
end;
$$ language plpgsql;
