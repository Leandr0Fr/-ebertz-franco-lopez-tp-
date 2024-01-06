create or replace function liquidar_obras_sociales(anio int, mes int) returns void as $$
declare
	l record;
	t record;
	datos_paciente record;
	datos_medique record;
	inicio_mes date;
	fin_mes date;
begin
	inicio_mes := date(anio::text || '-' || mes::text || '-' || '01');
	fin_mes := inicio_mes + '1 month'::interval - '1 day'::interval;

	for t in select nro_obra_social_consulta, sum(monto_obra_social) monto_total from turno 
	where date(fecha) between inicio_mes and fin_mes and estado = 'atendido' group by nro_obra_social_consulta loop	
		insert into liquidacion_cabecera values (default, t.nro_obra_social_consulta, inicio_mes, fin_mes, t.monto_total);
		
		update turno set estado = 'liquidado' where nro_obra_social_consulta = t.nro_obra_social_consulta and estado = 'atendido';
	end loop;
	
	for l in select * from liquidacion_cabecera where desde = inicio_mes and hasta = fin_mes loop
		for t in select * from turno where nro_obra_social_consulta = l.nro_obra_social and estado = 'liquidado' loop
			
			select * into datos_paciente from paciente where nro_paciente = t.nro_paciente;
			select * into datos_medique from medique where dni_medique = t.dni_medique;
		
			insert into liquidacion_detalle values (l.nro_liquidacion, default, t.fecha, t.nro_afiliade_consulta, datos_paciente.dni_paciente, 
				datos_paciente.nombre, datos_paciente.apellido, t.dni_medique, datos_medique.nombre, datos_medique.apellido, datos_medique.especialidad, t.monto_obra_social);
		end loop;
	end loop;
end;
$$ language plpgsql;
