create or replace function atender_paciente(num_turno int) returns boolean as $$
declare
	turno_selec record;
begin
	select * into turno_selec from turno where turno.nro_turno = num_turno;
	
	-- el numero de turno no existe
	if not found then
		insert into error values(default, null, null, null, null, 'atencion', current_timestamp, '?nro de turno no valido');
		return false;
	end if;
	
	-- turno no reservado
	if turno_selec.estado != 'reservado' then
		insert into error values(default, turno_selec.fecha, turno_selec.nro_consultorio,
		turno_selec.dni_medique, turno_selec.nro_paciente, 'atencion', current_timestamp, '?turno no reservado');
		return false;
	end if;
	
	-- el turno no corresponde con el dia actual
	if date(turno_selec.fecha) != current_date then
		insert into error values(default, turno_selec.fecha, turno_selec.nro_consultorio,
		turno_selec.dni_medique, turno_selec.nro_paciente, 'atencion', current_timestamp, '?turno no corresponde a la fecha del dia');
		return false;
	end if;
	
	-- se validaron todos los datos
	update turno set estado = 'atendido' where turno.nro_turno = num_turno;
	return true;
end;	
$$ language plpgsql;
