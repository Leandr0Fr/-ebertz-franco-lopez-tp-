create or replace function generar_turnos_disponibles(anio int, mes int) returns boolean as $$
declare
	fecha timestamp;
	a record;
	fecha_hora_ultimo_turno timestamp;
	fecha_hora_turno timestamp;
begin
	if existe_turno_para_mes(anio, mes) then
		return false;
	end if;
	
	for i in 1..cant_dias_de_mes(anio, mes) loop
		fecha := to_timestamp(anio::text || '-' || mes::text || '-' || i::text, 'yyyy-mmm-dd hh:mi:ss'); 
		
		for a in select * from agenda where nro_dia_de_semana(fecha) = dia loop
			fecha_hora_ultimo_turno := fecha + (a.hora_hasta - a.duracion_turno);
			fecha_hora_turno := fecha + a.hora_desde;
			
			while fecha_hora_turno <= fecha_hora_ultimo_turno loop
				insert into turno values(default, fecha_hora_turno, a.nro_consultorio, a.dni_medique,
					null, null, null, null, null, null, 'disponible');
				fecha_hora_turno = fecha_hora_turno + a.duracion_turno;
			end loop;
			
		end loop;
		
	end loop;
	
	return true;
end;
$$ language plpgsql;


create or replace function existe_turno_para_mes(anio int, mes int) returns boolean as $$
declare
	turno_buscado record;
begin
	select * into turno_buscado from turno where extract(month from fecha) = mes and extract(year from fecha) = anio;
	return found;
end;
$$ language plpgsql;


create or replace function nro_dia_de_semana(fecha timestamp) returns int as $$
declare
	num_dia int;
begin
	select extract(dow from fecha) into num_dia;
	return num_dia;
end;
$$ language plpgsql;
	
	
create or replace function cant_dias_de_mes(anio int, mes int) returns int as $$
declare
	cant_dias int;
	fecha_1 timestamp;
	fecha_2 timestamp;
begin
	fecha_1 := to_timestamp(anio::text || '-' || mes::text || '-01', 'yyyy-mmm-dd hh:mi:ss');
	fecha_2 := to_timestamp(anio::text || '-' || (mes + 1)::text || '-01', 'yyyy-mmm-dd hh:mi:ss');
	 
	select abs(extract(day from fecha_1 - fecha_2)) into cant_dias;
	return cant_dias;
end;
$$ language plpgsql;	
	
